package ApproxSS

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/tuneinsight/lattigo/v4/bfv"
	"github.com/tuneinsight/lattigo/v4/drlwe"
	"github.com/tuneinsight/lattigo/v4/ring"
	"github.com/tuneinsight/lattigo/v4/rlwe"
	"github.com/tuneinsight/lattigo/v4/utils"
)

type ExistingShamirApproxSS2 struct {
	VanSS *VanillaShamirSS

	params4cmb rlwe.Parameters
}

func NewExistingShamirApproxSS2(N, T int, params bfv.Parameters) (this *ExistingShamirApproxSS2) {
	this = new(ExistingShamirApproxSS2)
	this.VanSS = NewVanillaShamirSS(N, T, params)

	this.params4cmb, _ = rlwe.NewParameters(params.LogN(), params.Q(), nil, 0, params.HammingWeight(), params.Sigma(), params.RingType(), rlwe.NewScale(1), params.DefaultNTTFlag())
	return
	//this.params4cmb, _ = rlwe.NewParameters(params.LogN(), params.Q()[:1], nil, 0, params.HammingWeight(), params.Sigma(), params.RingType(), rlwe.NewScale(1), params.DefaultNTTFlag())

}

func (this *ExistingShamirApproxSS2) ApproxRecover(bound4smudgingNoise int) (isSuc bool, timeComp time.Duration, sizeComm float64) {

	paramsMesSpace := this.VanSS.thdizer.params
	ringQ := paramsMesSpace.RingQ()
	levelQ := 0

	//Round 1 Starts

	//Select the participants in round 1
	parR1 := SelectRandomParticipants(this.VanSS.N, this.VanSS.T)
	var timeStage1, timeStage2, timeStage3 time.Duration
	var sizeCommStage1, sizeCommStage2 float64
	//Parties run PartyR1
	for _, par := range parR1 {
		timeStage1Start := time.Now()
		prng_i, _ := utils.NewPRNG()
		smudgingNoiseSampler := ring.NewGaussianSampler(prng_i, paramsMesSpace.RingQ(), float64(bound4smudgingNoise)/6, bound4smudgingNoise)

		ni := ringQ.NewPoly()
		smudgingNoiseSampler.ReadLvl(levelQ, ni)
		timeStage1 = time.Since(timeStage1Start)
		_, cleanTimeStage1, sizeCommEachPiece := this.VanSS.ShareThenWrite(ni, "nonceShare", &par)
		//fmt.Println(cleanTimeStage1)
		timeStage1 = timeStage1 + cleanTimeStage1
		sizeCommStage1 = sizeCommEachPiece * float64(this.VanSS.N)

	}

	//Round 2 Starts

	//Select the participants in round 2
	parR2 := SelectRandomParticipants(this.VanSS.N, this.VanSS.T)
	//The output of parties in Round 2
	approxShareAll := make(map[int]*drlwe.ShamirSecretShare, this.VanSS.T)
	//Parties run PartyR2
	for _, par := range parR2 {
		//Read the share of secret from file
		filename := fmt.Sprintf("%s%d", "mesToParty", par)
		skFilename := fmt.Sprintf("%s%s", "file/skShare/", filename)
		skFile, _ := os.Open(skFilename)
		defer skFile.Close()
		skData, _ := io.ReadAll(skFile)
		share := this.VanSS.thdizer.AllocateThresholdSecretShare()
		share.UnmarshalBinary(skData)

		for _, i := range parR1 {
			nonce := this.VanSS.thdizer.AllocateThresholdSecretShare()
			nonceData := readNonceShare(par, i)
			nonce.UnmarshalBinary(nonceData)

			timeStage2Start := time.Now()
			this.VanSS.thdizer.AggregateShares(nonce, share, share)
			timeStage2 = time.Since(timeStage2Start)
		}

		approxShareAll[par] = share
		sizeCommStage2 = float64(len(skData))
		timeStage2 = timeStage2 * time.Duration(this.VanSS.T)
	}

	//AggregatorR2
	timeStage3Start := time.Now()
	points := make([]drlwe.ShamirPublicPoint, len(parR2))
	for idx, pIdx := range parR2 {
		points[idx] = drlwe.ShamirPublicPoint(pIdx + 1)
	}

	threshold := len(points)

	cmb := NewMyCombiner(&this.params4cmb, points, threshold)

	shares := make(map[drlwe.ShamirPublicPoint]drlwe.ShamirSecretShare)
	for i, share := range approxShareAll {
		pt := drlwe.ShamirPublicPoint(i + 1)
		shares[pt] = *share
	}
	result := cmb.Recover(shares)
	approxMessage := result.Q
	timeStage3 = time.Since(timeStage3Start)

	sizeComm = sizeCommStage1 + sizeCommStage2
	timeComp = timeStage1 + timeStage2 + timeStage3
	fmt.Println(timeStage1, timeStage2, timeStage3, timeComp)
	//Round 2 Ends!

	isSuc = TestResult(approxMessage, this.VanSS.secret.Value.Q, uint64(this.VanSS.T*bound4smudgingNoise), this.VanSS.thdizer.params.Q()[0:1])
	return
}

func (this *ExistingShamirApproxSS2) ApproxRecover4TestTime(bound4smudgingNoise int) (isSuc bool, timeComp time.Duration, sizeComm float64) {

	paramsMesSpace := this.VanSS.thdizer.params
	ringQ := paramsMesSpace.RingQ()
	levelQ := 0

	//Round 1 Starts

	//Select the participants in round 1
	parR1 := SelectRandomParticipants(this.VanSS.N, this.VanSS.T)
	var timeStage1, timeStage2, timeStage3 time.Duration
	var sizeCommStage1, sizeCommStage2 float64
	//Parties run PartyR1
	for i, par := range parR1 {
		if i == 0 {
			timeStage1Start := time.Now()
			prng_i, _ := utils.NewPRNG()
			smudgingNoiseSampler := ring.NewGaussianSampler(prng_i, paramsMesSpace.RingQ(), float64(bound4smudgingNoise)/6, bound4smudgingNoise)

			ni := ringQ.NewPoly()
			smudgingNoiseSampler.ReadLvl(levelQ, ni)
			timeStage1 = time.Since(timeStage1Start)
			_, cleanTimeStage1, sizeCommEachPiece := this.VanSS.ShareThenWrite(ni, "nonceShare", &par)
			timeStage1 = timeStage1 + cleanTimeStage1
			sizeCommStage1 = sizeCommEachPiece * float64(this.VanSS.N)
		}
	}

	//Round 2 Starts

	//Select the participants in round 2
	parR2 := SelectRandomParticipants(this.VanSS.N, this.VanSS.T)
	//The output of parties in Round 2
	approxShareAll := make(map[int]*drlwe.ShamirSecretShare, this.VanSS.T)
	//Parties run PartyR2
	for i, par := range parR2 {
		if i == 0 {
			// Read the share of secret from file
			filename := fmt.Sprintf("%s%d", "mesToParty", par)
			skFilename := fmt.Sprintf("%s%s", "file/skShare/", filename)
			skFile, _ := os.Open(skFilename)
			defer skFile.Close()
			skData, _ := io.ReadAll(skFile)
			share := this.VanSS.thdizer.AllocateThresholdSecretShare()
			share.UnmarshalBinary(skData)

			for _, j := range parR1 {
				if j == parR1[0] {
					nonce := this.VanSS.thdizer.AllocateThresholdSecretShare()
					nonceData := readNonceShare(par, j)
					nonce.UnmarshalBinary(nonceData)

					timeStage2Start := time.Now()
					this.VanSS.thdizer.AggregateShares(nonce, share, share)
					timeStage2 = time.Since(timeStage2Start)
				}
			}

			approxShareAll[par] = share
			sizeCommStage2 = float64(len(skData))
			timeStage2 = timeStage2 * time.Duration(this.VanSS.T)
		} else {
			approxShareAll[par] = approxShareAll[parR2[0]]
		}
	}

	//AggregatorR2
	timeStage3Start := time.Now()
	points := make([]drlwe.ShamirPublicPoint, len(parR2))
	for idx, pIdx := range parR2 {
		points[idx] = drlwe.ShamirPublicPoint(pIdx + 1)
	}

	threshold := len(points)

	cmb := NewMyCombiner(&this.params4cmb, points, threshold)

	shares := make(map[drlwe.ShamirPublicPoint]drlwe.ShamirSecretShare)
	for i, share := range approxShareAll {
		pt := drlwe.ShamirPublicPoint(i + 1)
		shares[pt] = *share
	}
	result := cmb.Recover(shares)
	approxMessage := result.Q
	timeStage3 = time.Since(timeStage3Start)

	sizeComm = sizeCommStage1 + sizeCommStage2
	timeComp = timeStage1 + timeStage2 + timeStage3
	fmt.Println(timeStage1, timeStage2, timeStage3, timeComp)
	//Round 2 Ends!

	isSuc = TestResult(approxMessage, this.VanSS.secret.Value.Q, uint64(this.VanSS.T*bound4smudgingNoise), this.VanSS.thdizer.params.Q()[0:1])
	return
}
