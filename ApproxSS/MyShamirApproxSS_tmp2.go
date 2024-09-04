package ApproxSS

// import (
// 	"fmt"
// 	"io"
// 	"os"
// 	"time"

// 	"github.com/didiercrunch/paillier"
// 	"github.com/tuneinsight/lattigo/v4/bfv"
// 	"github.com/tuneinsight/lattigo/v4/drlwe"
// 	"github.com/tuneinsight/lattigo/v4/ring"
// 	"github.com/tuneinsight/lattigo/v4/rlwe"
// 	"github.com/tuneinsight/lattigo/v4/utils"
// )

// type MyShamirApproxSS struct {
// 	VanSS *VanillaShamirSS

// 	AHEsks []*paillier.ThresholdPrivateKey
// 	AHEpk  paillier.PublicKey

// 	//params4cmb rlwe.Parameters
// 	//f          *skEncryptionBFV
// 	ek1All []*rlwe.SecretKey
// 	ek2All []*rlwe.SecretKey
// }

// func NewMyShamirApproxSS(N, T int, params bfv.Parameters) (myShamirApproxSS *MyShamirApproxSS) {
// 	myShamirApproxSS = new(MyShamirApproxSS)
// 	myShamirApproxSS.VanSS = NewVanillaShamirSS(N, T, params)

// 	return
// }

// func (myShamirApproxSS *MyShamirApproxSS) Preprocessing(bound4smudgingNoise int) (timeComp time.Duration, sizeComm float64) {

// 	paramsMesSpace := myShamirApproxSS.VanSS.thdizer.params
// 	ringQ := paramsMesSpace.RingQ()
// 	levelQ := 0

// 	// var timeStage1, timeStage2, timeStage3, timeStage4 time.Duration
// 	// var sizeCommStage1, sizeCommStage2, sizeCommStage3 float64

// 	//Round 1 Starts

// 	//Samples the polynomial a in our paper
// 	CRS, _ := utils.NewPRNG()
// 	a := TagGen(myShamirApproxSS.params4DoubleEncryption, CRS)

// 	//Parties run PartyR1
// 	for i := 0; i < myShamirApproxSS.VanSS.N; i++ {
// 		//Read the share from file
// 		filename := fmt.Sprintf("%s%d", "mesToParty", i)
// 		skFilename := fmt.Sprintf("%s%s", "file/skShare/", filename)
// 		skFile, _ := os.Open(skFilename)
// 		defer skFile.Close()
// 		skData, _ := io.ReadAll(skFile)
// 		share := myShamirApproxSS.VanSS.thdizer.AllocateThresholdSecretShare()
// 		share.UnmarshalBinary(skData)

// 		timeStage1Start := time.Now()
// 		prng_i, _ := utils.NewPRNG()
// 		smudgingNoiseSampler := ring.NewGaussianSampler(prng_i, paramsMesSpace.RingQ(), float64(bound4smudgingNoise)/6, bound4smudgingNoise)

// 		ni := ringQ.NewPoly()
// 		smudgingNoiseSampler.ReadLvl(levelQ, ni)
// 		CTni_all := make(map[int]*rlwe.Ciphertext, myShamirApproxSS.VanSS.T)

// 		CTni_all[i] = myShamirApproxSS.f.EncPolyCoeff(a, myShamirApproxSS.ek2All[i], ni)
// 		timeStage1 := time.Since(timeStage1Start)
// 		sizeCommStage1 := float64(CTni_all[i].Value[0].MarshalBinarySize64())
// 		fmt.Println(timeStage1, sizeCommStage1)
// 	}
// 	return
// }

// func (myShamirApproxSS *MyShamirApproxSS) ApproxRecover(bound4smudgingNoise int) (isSuc bool, timeComp time.Duration, sizeComm float64) {

// 	paramsMesSpace := myShamirApproxSS.VanSS.thdizer.params
// 	ringQ := paramsMesSpace.RingQ()
// 	levelQ := 0

// 	var timeStage1, timeStage2, timeStage3, timeStage4 time.Duration
// 	var sizeCommStage1, sizeCommStage2, sizeCommStage3 float64

// 	//Round 1 Starts

// 	//Samples the polynomial a in our paper
// 	CRS, _ := utils.NewPRNG()
// 	a := TagGen(myShamirApproxSS.params4DoubleEncryption, CRS)
// 	//Select the participants in round 1
// 	parR1 := SelectRandomParticipants(myShamirApproxSS.VanSS.N, myShamirApproxSS.VanSS.T)
// 	//The output of parties and aggregator in Round 1
// 	CTni_all := make(map[int]*rlwe.Ciphertext, myShamirApproxSS.VanSS.T)
// 	CTsi_all := make(map[int]*rlwe.Ciphertext, myShamirApproxSS.VanSS.T)
// 	lagrangeCoeffs := make(map[int]uint64, myShamirApproxSS.VanSS.T)

// 	//Parties run PartyR1
// 	for _, par := range parR1 {
// 		//Read the share from file
// 		filename := fmt.Sprintf("%s%d", "mesToParty", par)
// 		skFilename := fmt.Sprintf("%s%s", "file/skShare/", filename)
// 		skFile, _ := os.Open(skFilename)
// 		defer skFile.Close()
// 		skData, _ := io.ReadAll(skFile)
// 		share := myShamirApproxSS.VanSS.thdizer.AllocateThresholdSecretShare()
// 		share.UnmarshalBinary(skData)

// 		timeStage1Start := time.Now()
// 		prng_i, _ := utils.NewPRNG()
// 		smudgingNoiseSampler := ring.NewGaussianSampler(prng_i, paramsMesSpace.RingQ(), float64(bound4smudgingNoise)/6, bound4smudgingNoise)

// 		ni := ringQ.NewPoly()
// 		smudgingNoiseSampler.ReadLvl(levelQ, ni)

// 		CTsi_all[par] = myShamirApproxSS.f.EncPolyCoeff(a, myShamirApproxSS.ek1All[par], share.Poly.Q)
// 		CTni_all[par] = myShamirApproxSS.f.EncPolyCoeff(a, myShamirApproxSS.ek2All[par], ni)
// 		timeStage1 = time.Since(timeStage1Start)
// 		sizeCommStage1 = float64(CTsi_all[par].Value[0].MarshalBinarySize64()) * 2
// 	}

// 	//AggregatorR1
// 	timeStage2Start := time.Now()
// 	survivialPublicPoint := []drlwe.ShamirPublicPoint{}
// 	//Collect survivialClient
// 	for _, pi := range parR1 {
// 		survivialPublicPoint = append(survivialPublicPoint, drlwe.ShamirPublicPoint(pi+1))
// 	}
// 	cmbR1 := NewMyCombiner(&myShamirApproxSS.params4cmb, survivialPublicPoint, myShamirApproxSS.VanSS.T)
// 	for _, cIdx := range parR1 {
// 		lagrangeCoeffs[cIdx] = fromRNSScalarToInt(myShamirApproxSS.params4cmb.RingQP(), cmbR1.LagrangeCoeffs[drlwe.ShamirPublicPoint(cIdx+1)])
// 	}
// 	timeStage2 = time.Since(timeStage2Start)
// 	sizeCommStage2 = float64(8 * len(lagrangeCoeffs))
// 	//Round 1 Ends!

// 	//Round 2 Starts!

// 	//Select the participants in round 1
// 	parR2 := SelectRandomParticipants(myShamirApproxSS.VanSS.N, myShamirApproxSS.VanSS.T)
// 	//The output of parties in Round 1
// 	dkShare_All := make(map[int]drlwe.ShamirSecretShare, myShamirApproxSS.VanSS.T)

// 	//Parties run PartyR2
// 	for _, par := range parR2 {
// 		timeStage3Start := time.Now()
// 		dkShare_All[par] = myShamirApproxSS.f.GenerateDKShareFile(parR1, lagrangeCoeffs, par)
// 		timeStage3 = time.Since(timeStage3Start)
// 		sizeCommStage3 = float64(dkShare_All[par].Poly.Q.MarshalBinarySize64())
// 	}

// 	//AggregatorR2
// 	timeStage4Start := time.Now()
// 	DK := myShamirApproxSS.f.GenerateDK(parR2, dkShare_All)
// 	approxMessage := ringQ.NewPoly()
// 	myShamirApproxSS.f.FEDecFinal(DK, CTni_all, CTsi_all, lagrangeCoeffs, parR1, approxMessage)
// 	timeStage4 = time.Since(timeStage4Start)

// 	//Round 2 Ends!
// 	timeComp = timeStage1 + timeStage2 + timeStage3 + timeStage4
// 	sizeComm = sizeCommStage1 + sizeCommStage2 + sizeCommStage3
// 	isSuc = TestResult(approxMessage, myShamirApproxSS.VanSS.secret.Value.Q, uint64(myShamirApproxSS.VanSS.T*bound4smudgingNoise), myShamirApproxSS.VanSS.thdizer.params.Q()[0:1])
// 	return

// }
