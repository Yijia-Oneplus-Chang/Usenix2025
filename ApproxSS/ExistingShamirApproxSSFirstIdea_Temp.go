package ApproxSS

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"os"
	"time"
)

type ExistingShamirApproxSS1 struct {
	T       int
	N       int
	len     int
	modulus *big.Int
	secret  *big.Int // Although the secret should be a vector with len elements, we only consider one element, as the complexity is exactly linear
}

func NewExistingShamirApproxSS1(N, T, len int) (this *ExistingShamirApproxSS1) {
	this = new(ExistingShamirApproxSS1)
	this.secret = big.NewInt(1)
	this.N = N
	this.T = T
	this.len = len
	this.modulus = big.NewInt(1)
	Nfactorial := big.NewInt(1)
	Nfactorial.MulRange(1, int64(N))
	modulusLowerBound := big.NewInt(1)
	modulusLowerBound.Mul(modulusLowerBound, Nfactorial)
	modulusLowerBound.Mul(modulusLowerBound, Nfactorial)
	modulusLowerBound.Mul(modulusLowerBound, Nfactorial)
	//Suppose that each noise is upper bouned by 5 and 20 = 4*5
	modulusLowerBound.Mul(modulusLowerBound, big.NewInt(int64(N*20)))

	for i := 0; i < 2; i = i * 2 {
		if modulusLowerBound.ProbablyPrime(20) {
			this.modulus = modulusLowerBound
			break
		}
		modulusLowerBound.Add(modulusLowerBound, big.NewInt(1))
	}

	this.secret, _ = rand.Int(rand.Reader, this.modulus)

	return
}

func (SS1 *ExistingShamirApproxSS1) ShareThenWrite(mes *big.Int) {

	if mes == nil {
		mes = SS1.secret
	}
	points := make([]int64, SS1.N)
	for i := range points {
		points[i] = int64(i + 1)
	}

	shareWhole := BGW_encoding_BigInt(SS1.secret, points, SS1.T, SS1.modulus)
	for i := range points {
		filename := fmt.Sprintf("%s%d", "file/skShare_BigInt/Share", i)
		pieceData := shareWhole[i].Bytes()
		err1 := os.WriteFile(filename, pieceData, 0666)
		if err1 != nil {
			panic(err1)
		}
	}

	return
}

func (SS1 *ExistingShamirApproxSS1) Recover() (result *big.Int, isSuc bool) {

	parR := SelectRandomParticipants(SS1.N, SS1.T)
	points := make([]int64, SS1.T)
	for i := range points {
		points[i] = int64(parR[i] + 1)
	}
	lagrangeCoeff := make([]*big.Int, SS1.T)
	Inv_xjminusxicoeff := big.NewInt(1)

	for i := range lagrangeCoeff {
		lagrangeCoeff[i] = big.NewInt(1)
		for j, point := range points {
			if i != j {
				lagrangeCoeff[i].Mul(lagrangeCoeff[i], big.NewInt(point))
				Inv_xjminusxicoeff.ModInverse(big.NewInt(point-points[i]), SS1.modulus)
				lagrangeCoeff[i].Mul(lagrangeCoeff[i], Inv_xjminusxicoeff)
				lagrangeCoeff[i].Mod(lagrangeCoeff[i], SS1.modulus)
			}
		}
	}

	result = big.NewInt(0)
	share := big.NewInt(0)

	for i := range parR {
		filename := fmt.Sprintf("%s%d", "file/skShare_BigInt/Share", parR[i])
		file, _ := os.Open(filename)
		defer file.Close()
		shareData, _ := io.ReadAll(file)
		share.SetBytes(shareData)
		share.Mul(share, lagrangeCoeff[i])
		result.Add(result, share)
		result.Mod(result, SS1.modulus)
	}

	if result.Cmp(SS1.secret) == 0 {
		isSuc = true
	} else {
		isSuc = false
	}

	return
}

func (SS1 *ExistingShamirApproxSS1) ApproxRecover(bound4smudgingNoise int) (result *big.Int, isSuc bool, timeComp time.Duration, sizeComm float64) {

	parR := SelectRandomParticipants(SS1.N, SS1.T)
	points := make([]int64, SS1.T)
	for i := range points {
		points[i] = int64(parR[i] + 1)
	}
	lagrangeCoeff := make([]*big.Int, SS1.T)
	Inv_xjminusxicoeff := big.NewInt(1)

	for i := range lagrangeCoeff {
		lagrangeCoeff[i] = big.NewInt(1)
		for j, point := range points {
			if i != j {
				lagrangeCoeff[i].Mul(lagrangeCoeff[i], big.NewInt(point))
				Inv_xjminusxicoeff.ModInverse(big.NewInt(point-points[i]), SS1.modulus)
				lagrangeCoeff[i].Mul(lagrangeCoeff[i], Inv_xjminusxicoeff)
				lagrangeCoeff[i].Mod(lagrangeCoeff[i], SS1.modulus)
			}
		}
	}

	result = big.NewInt(0)
	share := big.NewInt(0)
	Nfactorial := big.NewInt(1)
	Nfactorial.MulRange(1, int64(SS1.N))

	for i := range parR {
		filename := fmt.Sprintf("%s%d", "file/skShare_BigInt/Share", parR[i])
		file, _ := os.Open(filename)
		defer file.Close()
		shareData, _ := io.ReadAll(file)
		share.SetBytes(shareData)
		//Simulate the process of each party
		random, _ := rand.Int(rand.Reader, big.NewInt(int64(bound4smudgingNoise)))

		random.Mul(random, Nfactorial)
		random.Mul(random, Nfactorial)
		share.Add(share, random)
		share.Mod(share, SS1.modulus)

		share.Mul(share, lagrangeCoeff[i])
		result.Add(result, share)
		result.Mod(result, SS1.modulus)
	}

	gap := big.NewInt(0)
	gap.Sub(result, SS1.secret)
	gap.Abs(gap)
	bound := big.NewInt(int64(bound4smudgingNoise * SS1.T))
	bound.Mul(bound, Nfactorial)
	bound.Mul(bound, Nfactorial)
	bound.Mul(bound, Nfactorial)
	if bound.Cmp(gap) == 1 {
		isSuc = true
	} else {
		isSuc = false
	}

	return
}

func (SS1 *ExistingShamirApproxSS1) ApproxRecover4TestTime(bound4smudgingNoise int) (result *big.Int, isSuc bool, timeComp time.Duration, sizeComm float64) {

	parR := SelectRandomParticipants(SS1.N, SS1.T)
	points := make([]int64, SS1.T)
	for i := range points {
		points[i] = int64(parR[i] + 1)
	}
	lagrangeCoeff := make([]*big.Int, SS1.T)
	Inv_xjminusxicoeff := big.NewInt(1)

	timeStartStage1 := time.Now()
	for i := range lagrangeCoeff {
		lagrangeCoeff[i] = big.NewInt(1)
		for j, point := range points {
			if i != j {
				lagrangeCoeff[i].Mul(lagrangeCoeff[i], big.NewInt(point))
				Inv_xjminusxicoeff.ModInverse(big.NewInt(point-points[i]), SS1.modulus)
				lagrangeCoeff[i].Mul(lagrangeCoeff[i], Inv_xjminusxicoeff)
				lagrangeCoeff[i].Mod(lagrangeCoeff[i], SS1.modulus)
			}
		}
	}

	result = big.NewInt(0)
	share := big.NewInt(0)
	Nfactorial := big.NewInt(1)
	Nfactorial.MulRange(1, int64(SS1.N))

	timeStage1 := time.Since(timeStartStage1)

	filename := fmt.Sprintf("%s%d", "file/skShare_BigInt/Share", parR[0])
	file, _ := os.Open(filename)
	defer file.Close()
	shareData, _ := io.ReadAll(file)
	share.SetBytes(shareData)

	//Simulate the process of each party
	timeStartStage2 := time.Now()
	random, _ := rand.Int(rand.Reader, big.NewInt(int64(bound4smudgingNoise)))

	random.Mul(random, Nfactorial)
	random.Mul(random, Nfactorial)
	share.Add(share, random)
	share.Mod(share, SS1.modulus)

	share.Mul(share, lagrangeCoeff[0])
	result.Add(result, share)
	result.Mod(result, SS1.modulus)

	timeStage2 := time.Since(timeStartStage2)

	timeComp = (timeStage1 + timeStage2) * time.Duration(SS1.len)

	sizeComm = float64(len(shareData)) * float64(SS1.len)

	isSuc = true

	return
}

// func (this *ExistingShamirApproxSS1) ApproxRecover(bound4smudgingNoise int) (isSuc bool) {

// 	paramsMesSpace := this.VanSS.thdizer.params
// 	ringQ := paramsMesSpace.RingQ()
// 	levelQ := 0

// 	//Start!
// 	var Nfactorial int
// 	Nfactorial = 1
// 	for i := 0; i < this.VanSS.N; i++ {
// 		Nfactorial = Nfactorial * i
// 	}
// 	NfactorialSquare := uint64(int(Nfactorial) * int(Nfactorial))
// 	NfactorialTriple := Nfactorial * int(NfactorialSquare)

// 	//Select the participants
// 	parR := SelectRandomParticipants(this.VanSS.N, this.VanSS.T)
// 	//The output of parties
// 	approxShareAll := make(map[int]*drlwe.ShamirSecretShare, this.VanSS.T)
// 	//Parties run PartyR2
// 	for _, par := range parR {
// 		//Read the share of secret from file
// 		filename := fmt.Sprintf("%s%d", "mesToParty", par)
// 		skFilename := fmt.Sprintf("%s%s", "file/skShare/", filename)
// 		skFile, _ := os.Open(skFilename)
// 		defer skFile.Close()
// 		skData, _ := io.ReadAll(skFile)
// 		share := this.VanSS.thdizer.AllocateThresholdSecretShare()
// 		share.UnmarshalBinary(skData)

// 		prng_i, _ := utils.NewPRNG()
// 		smudgingNoiseSampler := ring.NewGaussianSampler(prng_i, paramsMesSpace.RingQ(), float64(bound4smudgingNoise)/6, bound4smudgingNoise)
// 		ni := ringQ.NewPoly()
// 		smudgingNoiseSampler.ReadLvl(levelQ, ni)
// 		ringQ.MulScalar(ni, NfactorialSquare, ni)

// 		nonce := this.VanSS.thdizer.AllocateThresholdSecretShare()
// 		nonce.Q = ni.CopyNew()
// 		this.VanSS.thdizer.AggregateShares(nonce, share, share)
// 		approxShareAll[par] = share
// 	}

// 	//AggregatorR2
// 	points := make([]drlwe.ShamirPublicPoint, len(parR))
// 	for idx, pIdx := range parR {
// 		points[idx] = drlwe.ShamirPublicPoint(pIdx + 1)
// 	}

// 	threshold := len(points)

// 	cmb := NewMyCombiner(&this.params4cmb, points, threshold)

// 	shares := make(map[drlwe.ShamirPublicPoint]drlwe.ShamirSecretShare)
// 	for i, share := range approxShareAll {
// 		pt := drlwe.ShamirPublicPoint(i + 1)
// 		shares[pt] = *share
// 	}
// 	result := cmb.Recover(shares)
// 	approxMessage := result.Q

// 	//End!

// 	isSuc = TestResult(approxMessage, this.VanSS.secret.Value.Q, uint64(this.VanSS.T*bound4smudgingNoise*int(NfactorialTriple)), this.VanSS.thdizer.params.Q()[0:1])
// 	return
// }
