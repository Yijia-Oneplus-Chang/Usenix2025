package ApproxSS

import (
	Crand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"

	"github.com/tuneinsight/lattigo/v4/drlwe"
	"github.com/tuneinsight/lattigo/v4/ring"
	"github.com/tuneinsight/lattigo/v4/rlwe"
	"github.com/tuneinsight/lattigo/v4/rlwe/ringqp"
	"github.com/tuneinsight/lattigo/v4/utils"
)

type MyThresholdizer struct {
	params   *rlwe.Parameters
	ringQP   *ringqp.Ring
	usampler ringqp.UniformSampler
	*drlwe.Thresholdizer
}

func BGW_encoding(X uint64, points []uint64, threshold int, p uint64) (X_BGW []uint64) {

	N := len(points)

	X_BGW = make([]uint64, N)
	for i := range X_BGW {
		X_BGW[i] = 0
	}

	R := make([]uint64, threshold)
	for i := range R {
		R[i] = rand.Uint64() % p
	}

	R[0] = X

	for i := 0; i < N; i++ {
		pointsExp := uint64(1)
		for t := 0; t < threshold; t++ {
			//X_BGW[i] = (X_BGW[i] + R[t]*ModExp(points[i], uint64(t), p)) % p
			X_BGW[i] = (X_BGW[i] + R[t]*pointsExp) % p
			pointsExp = pointsExp * points[i]
			if pointsExp >= p {
				pointsExp = pointsExp % p
			}
		}
	}

	return
}

func BGW_encoding_BigInt(X *big.Int, points []int64, threshold int, p *big.Int) (X_BGW []*big.Int) {

	N := len(points)

	X_BGW = make([]*big.Int, N)
	for i := range X_BGW {
		X_BGW[i] = big.NewInt(0)
	}

	R := make([]*big.Int, threshold)
	for i := range R {
		R[i], _ = Crand.Int(Crand.Reader, p)
	}

	R[0] = X

	for i := 0; i < N; i++ {
		pointsExp := big.NewInt(1)
		for t := 0; t < threshold; t++ {
			//X_BGW[i] = (X_BGW[i] + R[t]*ModExp(points[i], uint64(t), p)) % p
			Rtemp := big.NewInt(1)
			X_BGW[i].Add(X_BGW[i], Rtemp.Mul(R[t], pointsExp)) // = (X_BGW[i] + R[t]*pointsExp) % p
			X_BGW[i].Mod(X_BGW[i], p)
			pointsExp.Mul(pointsExp, big.NewInt(points[i])) // pointsExp = pointsExp * points[i]
			pointsExp.Mod(pointsExp, p)
		}
	}

	return
}

func NewMyThresholdizer(params rlwe.Parameters) *MyThresholdizer {

	Thr := new(MyThresholdizer)
	Thr.Thresholdizer = drlwe.NewThresholdizer(params)
	Thr.params = &params
	Thr.ringQP = params.RingQP()

	prng, err := utils.NewPRNG()
	if err != nil {
		panic(fmt.Errorf("could not initialize PRNG: %s", err))
	}

	Thr.usampler = ringqp.NewUniformSampler(prng, *params.RingQP())

	return Thr
}

// GenShamirSecretShares generates a secret share for the given recipient, identified by its ShamirPublicPoint.
// The result is stored in ShareOut and should be sent to this party.
func (Thr *MyThresholdizer) GenShamirSecretShares(threshold int, recipients []drlwe.ShamirPublicPoint, secret *rlwe.SecretKey) (sharesOut []*drlwe.ShamirSecretShare) {
	Qmodulus := Thr.ringQP.RingQ.Modulus
	var Pmodulus []uint64
	if Thr.ringQP.RingP != nil {
		Pmodulus = Thr.ringQP.RingP.Modulus
	}

	// coeffsDataQ := make([]float64, len(secret.Value.Q.Buff))
	// for i, coeff := range secret.Value.Q.Buff {
	// 	coeffsDataQ[i] = float64(coeff)
	// }

	points := make([]uint64, len(recipients))
	for i := range points {
		points[i] = uint64(recipients[i])
	}

	for i := 0; i < len(recipients); i++ {
		shareOut := Thr.AllocateThresholdSecretShare()
		sharesOut = append(sharesOut, shareOut)
	}

	for qdx, q := range Qmodulus {
		for cdx, coeff := range secret.Value.Q.Coeffs[qdx] {
			shares := BGW_encoding(coeff, points, threshold, q)
			for idx, share := range shares {
				sharesOut[idx].Q.Coeffs[qdx][cdx] = share
			}
		}
	}

	for _, shareOut := range sharesOut {
		shareOut.Q.Buff = []uint64{}
		for _, coeffs := range shareOut.Q.Coeffs {
			shareOut.Q.Buff = append(shareOut.Q.Buff, coeffs...)
		}
	}

	if Thr.ringQP.RingP != nil {
		for pdx, p := range Pmodulus {
			for cdx, coeff := range secret.Value.P.Coeffs[pdx] {
				shares := BGW_encoding(coeff, points, threshold, p)
				for idx, share := range shares {
					sharesOut[idx].P.Coeffs[pdx][cdx] = share
				}
			}
		}

		for _, shareOut := range sharesOut {
			shareOut.P.Buff = []uint64{}
			for _, coeffs := range shareOut.P.Coeffs {
				shareOut.P.Buff = append(shareOut.P.Buff, coeffs...)
			}
		}
	}

	return
}

// // GenShamirSecretShares generates a secret share for the given recipient, identified by its ShamirPublicPoint.
// // The result is stored in ShareOut and should be sent to this party.
// func (Thr *MyThresholdizer) GenShamirSecretSharesComparedWith(threshold int, recipients []drlwe.ShamirPublicPoint, secret *rlwe.SecretKey) (sharesOut []drlwe.ShamirSecretShare) {

// 	poly, _ := Thr.GenShamirPolynomial(threshold, secret)

// 	for _, recipient := range recipients {
// 		shareOut := Thr.AllocateThresholdSecretShare()
// 		Thr.GenShamirSecretShare(recipient, poly, shareOut)
// 		sharesOut = append(sharesOut, *shareOut)
// 	}

// 	return
// }

type MyCombiner struct {
	params *rlwe.Parameters
	ringQP *ringqp.Ring
	//usampler  ringqp.UniformSampler
	threshold  int
	tmp1, tmp2 []uint64
	one        ring.RNSScalar

	LagrangeCoeffs map[drlwe.ShamirPublicPoint]ring.RNSScalar
	//LagrangeCoeffsGlobal map[drlwe.ShamirPublicPoint]map[drlwe.ShamirPublicPoint]ring.RNSScalar
	points []drlwe.ShamirPublicPoint
}

func NewMyCombiner(params *rlwe.Parameters, points []drlwe.ShamirPublicPoint, threshold int) *MyCombiner {
	cmb := new(MyCombiner)
	cmb.params = params
	cmb.ringQP = params.RingQP()
	cmb.threshold = threshold
	cmb.tmp1, cmb.tmp2 = cmb.ringQP.NewRNSScalar(), cmb.ringQP.NewRNSScalar()
	cmb.one = ScalarTransform(cmb.ringQP, cmb.ringQP.NewRNSScalarFromUInt64(1))
	cmb.points = points

	cmb.LagrangeCoeffs = make(map[drlwe.ShamirPublicPoint]ring.RNSScalar, threshold)
	//cmb.LagrangeCoeffsGlobal = make(map[drlwe.ShamirPublicPoint]map[drlwe.ShamirPublicPoint]ring.RNSScalar)

	prod := cmb.tmp2

	for _, thisPoint := range points {
		cmb.LagrangeCoeffs[thisPoint] = cmb.ringQP.NewRNSScalar()
		copy(prod, cmb.one)
		for _, thatPoint := range points {
			//Lagrange Interpolation with the public threshold key of other active players
			if thisPoint != thatPoint {
				cmb.lagrangeCoeff(thisPoint, thatPoint, cmb.tmp1)
				cmb.ringQP.MulRNSScalar(prod, cmb.tmp1, prod)
			}
		}
		copy(cmb.LagrangeCoeffs[thisPoint], prod)
	}

	return cmb
}

func (cmb *MyCombiner) lagrangeCoeff(thisKey drlwe.ShamirPublicPoint, thatKey drlwe.ShamirPublicPoint, lagCoeff []uint64) {

	this := cmb.ringQP.NewRNSScalarFromUInt64(uint64(thisKey))
	that := cmb.ringQP.NewRNSScalarFromUInt64(uint64(thatKey))

	cmb.ringQP.SubRNSScalar(that, this, lagCoeff)

	cmb.ringQP.Inverse(lagCoeff)

	cmb.ringQP.MulRNSScalar(lagCoeff, that, lagCoeff)
}

func (cmb *MyCombiner) Recover(shares map[drlwe.ShamirPublicPoint]drlwe.ShamirSecretShare) *ringqp.Poly {
	if len(shares) != cmb.threshold {
		fmt.Println(len(shares))
		fmt.Println(cmb.threshold)
		panic("The number of shares is not equal to the threshold!")
	}

	result := cmb.params.RingQP().NewPoly()

	for point, share := range shares {
		//point := cmb.points[idx]
		resultTemp := cmb.params.RingQP().NewPoly()
		cmb.ringQP.MulRNSScalarMontgomery(share.Poly, cmb.LagrangeCoeffs[point], resultTemp)
		cmb.ringQP.AddLvl(cmb.params.QCount()-1, cmb.params.PCount()-1, resultTemp, result, result)

	}

	return &result
}

func (cmb *MyCombiner) RecoverQLvl(shares map[drlwe.ShamirPublicPoint]drlwe.ShamirSecretShare, level int) *ring.Poly {
	if len(shares) != cmb.threshold {
		fmt.Println(len(shares))
		fmt.Println(cmb.threshold)
		panic("The number of shares is not equal to the threshold!")
	}

	result := cmb.params.RingQ().NewPolyLvl(level)
	LagCoeffsLvl := make(map[drlwe.ShamirPublicPoint]ring.RNSScalar, len(cmb.LagrangeCoeffs))

	//scalarTmp := ScalarTransform(S.f.bfvParams.RingQP(), S.params.RingQ().NewRNSScalarFromUInt64())
	//S.params.RingQ().MulRNSScalarMontgomeryLvl(0, noise, scalarTmp, noiseTmp)
	//S.params.RingQ().AddLvl(0, noiseResultWant, noiseTmp, noiseResultWant)

	for i, li := range cmb.LagrangeCoeffs {
		LagCoeffsLvl[i] = cmb.ringQP.RingQ.NewRNSScalarFromUInt64(fromRNSScalarToInt(cmb.ringQP, li))
		LagCoeffsLvl[i] = ScalarTransformQLvl(cmb.ringQP.RingQ, LagCoeffsLvl[i], level)
	}

	for point, share := range shares {
		//point := cmb.points[idx]
		resultTemp := cmb.params.RingQ().NewPolyLvl(level)
		cmb.ringQP.RingQ.MulRNSScalarMontgomeryLvl(level, share.Poly.Q, LagCoeffsLvl[point], resultTemp)
		cmb.ringQP.RingQ.AddLvl(level, resultTemp, result, result)

	}

	return result
}

// func testCombiner() {
// 	params, _ := rlwe.NewParametersFromLiteral(rlwe.TestPN11QP54)
// 	numClient := 16
// 	threshold := 12
// 	points := make([]drlwe.ShamirPublicPoint, numClient)
// 	for i := 0; i < numClient; i++ {
// 		points[i] = drlwe.ShamirPublicPoint(i + 1)
// 	}

// 	cmb := NewMyCombiner(&params, points[:threshold], threshold)

// 	keyGen := rlwe.NewKeyGenerator(params)
// 	sk := keyGen.GenSecretKey()
// 	thr := drlwe.NewThresholdizer(params)
// 	skPoly, _ := thr.GenShamirPolynomial(threshold, sk)
// 	shares := make([]drlwe.ShamirSecretShare, threshold)
// 	for i := 0; i < threshold; i++ {
// 		ShareTemp := thr.AllocateThresholdSecretShare()
// 		thr.GenShamirSecretShare(points[i], skPoly, ShareTemp)
// 		shares[i] = *ShareTemp
// 	}

// 	result := cmb.Recover(shares[:threshold])
// 	fmt.Println(result.Equals(sk.Value))
// }
