package ApproxSS

import (
	"fmt"
	"time"

	"github.com/tuneinsight/lattigo/v4/bfv"
	"github.com/tuneinsight/lattigo/v4/drlwe"
	"github.com/tuneinsight/lattigo/v4/ring"
	"github.com/tuneinsight/lattigo/v4/rlwe"
	"github.com/tuneinsight/lattigo/v4/rlwe/ringqp"
	"github.com/tuneinsight/lattigo/v4/utils"
)

func TagGen(params bfv.Parameters, prng *utils.KeyedPRNG) *ring.Poly {
	tag := params.RingQP().NewPoly()
	uniformSampler := ringqp.NewUniformSampler(prng, *params.RingQP())
	uniformSampler.ReadLvl(params.MaxLevel(), -1, tag)
	return tag.Q
}

type skEncryptionBFV struct {
	bfvParams           bfv.Parameters
	gaussianSamplerQ4FE *ring.GaussianSampler
	uniformSamplerQ4FE  ringqp.UniformSampler
	thr4FE              *MyThresholdizer
	encoder             bfv.Encoder
}

// func NewFE(params rlwe.Parameters, T uint64) *skEncryptionBFV {
// 	f := new(FE)
// 	f.params = params
// 	f.T = T
// 	prng, _ := utils.NewPRNG()
// 	f.gaussianSamplerQ4FE = ring.NewGaussianSampler(prng, f.params.RingQ(), f.params.Sigma(), int(6*f.params.Sigma()))
// 	f.uniformSamplerQ4FE = ring.NewUniformSampler(prng, f.params.RingQ())
// 	f.thr4FE = drlwe.NewThresholdizer(f.params)
// 	deltaInt := f.params.Q()[0] / f.T
// 	*f.delta = f.params.RingQ().NewRNSScalarFromUInt64(deltaInt)
// 	f.scaler = bfv.NewRNSScaler(f.params.RingQ(), f.T)

// 	return f
// }

func NewSKencryption(params bfv.Parameters) (f *skEncryptionBFV) {

	f = new(skEncryptionBFV)
	f.bfvParams = params

	prng, _ := utils.NewPRNG()
	f.gaussianSamplerQ4FE = ring.NewGaussianSampler(prng, f.bfvParams.RingQ(), f.bfvParams.Sigma(), int(6*f.bfvParams.Sigma()))
	f.uniformSamplerQ4FE = ringqp.NewUniformSampler(prng, *f.bfvParams.RingQP())
	f.thr4FE = NewMyThresholdizer(f.bfvParams.Parameters)
	f.encoder = bfv.NewEncoder(f.bfvParams)
	return
}

func (f *skEncryptionBFV) EncPolyCoeff(tag *ring.Poly, sk *rlwe.SecretKey, mes *ring.Poly) *rlwe.Ciphertext {

	mesVector := make([]uint64, mes.N())
	for i := 0; i < mes.N(); i++ {
		mesVector[i] = mes.Buff[i]
	}

	ct := f.FEEnc(tag, sk, mesVector)
	return ct
}

func (f *skEncryptionBFV) FEEnc(tag *ring.Poly, sk *rlwe.SecretKey, mesVector []uint64) *rlwe.Ciphertext {

	pt := bfv.NewPlaintext(f.bfvParams, f.bfvParams.MaxLevel())
	f.encoder.Encode(mesVector, pt)

	ct := bfv.NewCiphertext(f.bfvParams, 1, pt.Level())
	ct.Value[1] = tag.CopyNew()
	c1 := ct.Value[1]

	ringQ := f.bfvParams.RingQ()
	levelQ := ct.Level()

	c0 := ct.Value[0]

	ringQ.MulCoeffsMontgomeryLvl(levelQ, c1, sk.Value.Q, c0) // c0 = NTT(sc1)
	ringQ.NegLvl(levelQ, c0, c0)                             // c0 = NTT(-sc1)

	buff := f.bfvParams.RingQ().NewPoly()
	if ct.IsNTT {
		f.gaussianSamplerQ4FE.ReadLvl(levelQ, buff) // e

		ringQ.NTTLvl(levelQ, buff, buff)   // NTT(e)
		ringQ.AddLvl(levelQ, c0, buff, c0) // c0 = NTT(-sc1 + e)
	} else {
		ringQ.InvNTTLvl(levelQ, c0, c0) // c0 = -sc1
		if ct.Degree() == 1 {
			ringQ.InvNTTLvl(levelQ, c1, c1) // c1 = c1
		}
		f.gaussianSamplerQ4FE.ReadAndAddLvl(levelQ, c0) // c0 = -sc1 + e
	}

	f.bfvParams.RingQ().AddLvl(ct.Level(), ct.Value[0], pt.Value, ct.Value[0])

	return ct

}

func (f *skEncryptionBFV) GenerateEKfromShareFile(N, generator int, NorS string) (sk *rlwe.SecretKey) {

	sk = rlwe.NewSecretKey(f.bfvParams.Parameters)
	levelP := sk.LevelP()
	levelQ := sk.LevelQ()
	uShare := f.thr4FE.AllocateThresholdSecretShare()
	for j := 0; j < N; j++ {
		uShare.UnmarshalBinary(readSK4FEShare(generator, j, NorS))
		f.bfvParams.Parameters.RingQP().AddLvl(levelQ, levelP, sk.Value, uShare.Poly, sk.Value)
	}

	return sk
}

func (f *skEncryptionBFV) GenerateEKThenShareToFile(N, T, generator int, NorS string) (*rlwe.SecretKey, time.Duration) {

	start := time.Now()
	keyGen4FE := bfv.NewKeyGenerator(f.bfvParams)
	sk := keyGen4FE.GenSecretKey()
	points := make([]drlwe.ShamirPublicPoint, N)
	for i := 0; i < N; i++ {
		points[i] = drlwe.ShamirPublicPoint(i + 1)
	}

	sk4FE1Share := f.thr4FE.GenShamirSecretShares(T, points, sk)

	timeTrue := time.Since(start)

	for i := 0; i < N; i++ {
		writeSK4FEShare(i, generator, NorS, *sk4FE1Share[i])
	}

	return sk, timeTrue
}

func (f *skEncryptionBFV) GenerateDKShareNew(N int, eknShare []*drlwe.ShamirSecretShare, eks *drlwe.ShamirSecretShare) drlwe.ShamirSecretShare {

	dkShare := f.thr4FE.AllocateThresholdSecretShare()

	dkShare.Poly = f.bfvParams.RingQP().NewPoly()

	for i := 0; i < N; i++ {
		f.thr4FE.AggregateShares(dkShare, eknShare[i], dkShare)
	}

	f.thr4FE.AggregateShares(dkShare, eks, dkShare)

	return *dkShare

}

func (f *skEncryptionBFV) GenerateDKShare(client4Encryption []int, yInt map[int]uint64, tsk []*drlwe.ShamirSecretShare) drlwe.ShamirSecretShare {
	yIn := make(map[int]ring.RNSScalar, len(yInt))
	for i, yi := range yInt {
		yIn[i] = f.bfvParams.RingQP().NewRNSScalarFromUInt64(yi)
	}
	dkShare := f.thr4FE.AllocateThresholdSecretShare()

	dkShare.Poly = f.bfvParams.RingQP().NewPoly()

	resultTemp := f.thr4FE.AllocateThresholdSecretShare()
	for _, pIdx := range client4Encryption {

		f.bfvParams.RingQ().MulRNSScalarMontgomery(tsk[pIdx].Q, ScalarTransform(f.bfvParams.RingQP(), yIn[pIdx]), resultTemp.Q)
		f.thr4FE.AggregateShares(dkShare, resultTemp, dkShare)
	}

	return *dkShare

}

func (f *skEncryptionBFV) GenerateDKShareFile(client4Encryption []int, yInt map[int]uint64, id int) (drlwe.ShamirSecretShare, time.Duration) {

	start := time.Now()
	yIn := make(map[int]ring.RNSScalar, len(yInt))
	for i, yi := range yInt {
		yIn[i] = f.bfvParams.RingQP().NewRNSScalarFromUInt64(yi)
	}

	dkShare := f.thr4FE.AllocateThresholdSecretShare()

	dkShare.Poly = f.bfvParams.RingQP().NewPoly()

	resultTemp := f.thr4FE.AllocateThresholdSecretShare()

	tempEks := f.thr4FE.AllocateThresholdSecretShare()
	tempEkn := f.thr4FE.AllocateThresholdSecretShare()

	for _, pIdx := range client4Encryption {
		tempEkn.UnmarshalBinary(readSK4FEShare(id, pIdx, "n"))
		tempEks.UnmarshalBinary(readSK4FEShare(id, pIdx, "s"))
		f.bfvParams.RingQ().MulRNSScalarMontgomery(tempEks.Q, ScalarTransform(f.bfvParams.RingQP(), yIn[pIdx]), resultTemp.Q)
		f.thr4FE.AggregateShares(dkShare, resultTemp, dkShare)
		f.thr4FE.AggregateShares(dkShare, tempEkn, dkShare)
	}

	timeTotal := time.Since(start)

	start4ReadFiles := time.Now()
	readSK4FEShare(id, client4Encryption[0], "n")
	time4ReadFiles := time.Since(start4ReadFiles)
	timeClean := timeTotal - time4ReadFiles*time.Duration(2*len(client4Encryption))
	return *dkShare, timeClean

}

func (f *skEncryptionBFV) GenerateDKShareFile4TestTime(client4Encryption []int, yInt map[int]uint64, id int) (drlwe.ShamirSecretShare, time.Duration) {

	start := time.Now()
	yIn := make(map[int]ring.RNSScalar, len(yInt))
	for i, yi := range yInt {
		yIn[i] = f.bfvParams.RingQP().NewRNSScalarFromUInt64(yi)
	}

	dkShare := f.thr4FE.AllocateThresholdSecretShare()

	dkShare.Poly = f.bfvParams.RingQP().NewPoly()

	resultTemp := f.thr4FE.AllocateThresholdSecretShare()

	tempEks := f.thr4FE.AllocateThresholdSecretShare()
	tempEkn := f.thr4FE.AllocateThresholdSecretShare()

	var time4ReadFiles time.Duration
	for i, pIdx := range client4Encryption {
		if i == 0 {
			start4ReadFiles := time.Now()
			tempEkn.UnmarshalBinary(readSK4FEShare(id, pIdx, "n"))
			tempEks.UnmarshalBinary(readSK4FEShare(id, pIdx, "s"))
			time4ReadFiles = time.Since(start4ReadFiles)
		}
		f.bfvParams.RingQ().MulRNSScalarMontgomery(tempEks.Q, ScalarTransform(f.bfvParams.RingQP(), yIn[pIdx]), resultTemp.Q)
		f.thr4FE.AggregateShares(dkShare, resultTemp, dkShare)
		f.thr4FE.AggregateShares(dkShare, tempEkn, dkShare)
	}

	timeTotal := time.Since(start)
	timeClean := timeTotal - time4ReadFiles
	return *dkShare, timeClean

}

// client4Decryption should be one-to-one correspondance to shares
// the i-th element of shares is the share of client4Decryption[i]
func (f *skEncryptionBFV) GenerateDK(client4Decryption []int, sharesInput map[int]drlwe.ShamirSecretShare) *rlwe.SecretKey {

	points := make([]drlwe.ShamirPublicPoint, len(client4Decryption))
	for idx, pIdx := range client4Decryption {
		points[idx] = drlwe.ShamirPublicPoint(pIdx + 1)
	}

	threshold := len(points)

	cmb := NewMyCombiner(&f.bfvParams.Parameters, points, threshold)

	shares := make(map[drlwe.ShamirPublicPoint]drlwe.ShamirSecretShare)
	for i, share := range sharesInput {
		pt := drlwe.ShamirPublicPoint(i + 1)
		shares[pt] = share
	}
	result := cmb.Recover(shares)
	dk := rlwe.NewSecretKey(f.bfvParams.Parameters)
	dk.Value = result.CopyNew()
	return dk
	//fmt.Println(result.Equals(sk.Value))
}

// func (f *skEncryptionBFV) FEDecFinalNew(dk *rlwe.SecretKey, cipherNoiseAll map[int]*rlwe.Ciphertext, cipherShareAll map[int]*rlwe.Ciphertext, yInt map[int]uint64, client4Encryption []int, result *ring.Poly) {

// 	if len(cipherShareAll) != len(yInt) {
// 		panic("The length of cipherShareAll and yInt are not match!")
// 	}

// 	y := make(map[int]ring.RNSScalar, len(yInt)+len(cipherNoiseAll))
// 	cipherAll := make(map[int]*rlwe.Ciphertext, len(yInt)+len(cipherNoiseAll))
// 	client4EncryptionExpand := make([]int, len(yInt)+len(cipherNoiseAll))

// 	count := int(0)
// 	for i, yi := range yInt {
// 		y[i] = f.bfvParams.RingQP().NewRNSScalarFromUInt64(yi)
// 		cipherAll[i] = cipherShareAll[i]
// 		client4EncryptionExpand[count] = i
// 		count = count + 1
// 	}

// 	for i, _ := range cipherNoiseAll {
// 		y[-1*i] = f.bfvParams.RingQP().NewRNSScalarFromUInt64(1)
// 		cipherAll[-1*i] = cipherNoiseAll[i]
// 		client4EncryptionExpand[count] = -1 * i
// 		count = count + 1
// 	}

// 	mesVector := f.FEDec(dk, cipherAll, y, client4EncryptionExpand)

// 	if len(result.Coeffs[0]) != len(mesVector) {
// 		panic("The length of result noise and decrypted vector are not match!")
// 	}

// 	result.Zero()

// 	for i, mi := range mesVector {
// 		result.Buff[i] = mi
// 		result.Coeffs[0][i] = mi
// 	}

//		return
//	}

func (f *skEncryptionBFV) FEDecFinalNew(dk *rlwe.SecretKey, cipherNoiseAll map[int]*rlwe.Ciphertext, cipherShareAll map[int]*rlwe.Ciphertext, yInt map[int]uint64, client4Encryption []int, result *ring.Poly) {

	if len(cipherShareAll) != len(yInt) {
		panic("The length of cipherShareAll and yInt are not match!")
	}

	y := make(map[int]ring.RNSScalar, len(yInt))
	//cipherAll := make(map[int]*rlwe.Ciphertext, len(yInt))
	//client4EncryptionExpand := make([]int, len(yInt))

	count := int(0)
	for i, yi := range yInt {
		y[i] = f.bfvParams.RingQP().NewRNSScalarFromUInt64(yi)
		//cipherAll[i] = cipherShareAll[i]
		//client4EncryptionExpand[count] = i
		count = count + 1
	}

	mesVector := f.FEDec(dk, cipherShareAll, y, client4Encryption)

	if len(result.Coeffs[0]) != len(mesVector) {
		panic("The length of result noise and decrypted vector are not match!")
	}

	result.Zero()

	for i, mi := range mesVector {
		result.Buff[i] = mi
		result.Coeffs[0][i] = mi
	}

	return
}

func (f *skEncryptionBFV) FEDecFinal(dk *rlwe.SecretKey, cipherNoiseAll map[int]*rlwe.Ciphertext, cipherShareAll map[int]*rlwe.Ciphertext, yInt map[int]uint64, client4Encryption []int, result *ring.Poly) {

	y := make(map[int]ring.RNSScalar, len(yInt)*2)
	for i, yi := range yInt {
		y[2*i] = f.bfvParams.RingQP().NewRNSScalarFromUInt64(yi)
		y[2*i+1] = f.bfvParams.RingQP().NewRNSScalarFromUInt64(1)
	}

	cipherAll := make(map[int]*rlwe.Ciphertext, len(cipherNoiseAll)*2)
	for i, _ := range cipherNoiseAll {
		cipherAll[2*i] = cipherShareAll[i]
		cipherAll[2*i+1] = cipherNoiseAll[i]
	}

	client4EncryptionDoubleExpand := make([]int, len(client4Encryption)*2)
	for i, ci := range client4Encryption {
		client4EncryptionDoubleExpand[2*i] = ci * 2
		client4EncryptionDoubleExpand[2*i+1] = ci*2 + 1
	}

	mesVector := f.FEDec(dk, cipherAll, y, client4EncryptionDoubleExpand)

	if len(result.Coeffs[0]) != len(mesVector) {
		panic("The length of result noise and decrypted vector are not match!")
	}

	result.Zero()

	for i, mi := range mesVector {
		result.Buff[i] = mi
		result.Coeffs[0][i] = mi
	}

	return
}

func (f *skEncryptionBFV) FEDec(dk *rlwe.SecretKey, cipherAll map[int]*rlwe.Ciphertext, yIn map[int]ring.RNSScalar, client4Encryption []int) []uint64 {

	ringQ := f.bfvParams.RingQ()

	ct := bfv.NewCiphertext(f.bfvParams, 1, f.bfvParams.MaxLevel())

	ct.Value[1] = cipherAll[client4Encryption[0]].Value[1].CopyNew()
	ct.Value[0] = cipherAll[client4Encryption[0]].Value[0].CopyNew()

	ringQ.MulRNSScalarMontgomery(ct.Value[0], ScalarTransform(f.bfvParams.RingQP(), yIn[client4Encryption[0]]), ct.Value[0])

	ciphertextTemp := ringQ.NewPoly()

	for i, ci := range client4Encryption {
		if i != 0 {
			if cipherAll[ci].Value[1].Equals(ct.Value[1]) {
				ringQ.MulRNSScalarMontgomery(cipherAll[ci].Value[0], ScalarTransform(f.bfvParams.RingQP(), yIn[ci]), ciphertextTemp)
				ringQ.Add(ct.Value[0], ciphertextTemp, ct.Value[0])
			} else {
				fmt.Println("Wrong ciphertexts for FE: dismatch of tag polynomial for ciphertext", ci)
			}
		}
	}

	decryptor := bfv.NewDecryptor(f.bfvParams, dk)
	pt := decryptor.DecryptNew(ct)

	result := f.encoder.DecodeUintNew(pt)

	return result

}
