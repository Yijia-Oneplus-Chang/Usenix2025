package ApproxSS

// import (
// 	"fmt"
// 	"io"
// 	"os"

// 	"github.com/tuneinsight/lattigo/v4/bfv"
// 	"github.com/tuneinsight/lattigo/v4/drlwe"
// 	"github.com/tuneinsight/lattigo/v4/ring"
// 	"github.com/tuneinsight/lattigo/v4/rlwe"
// 	"github.com/tuneinsight/lattigo/v4/utils"
// )

// type ExistingShamirApproxSS1 struct {
// 	VanSS *VanillaShamirSS

// 	params4cmb rlwe.Parameters
// }

// func NewExistingShamirApproxSS1(N, T int, params bfv.Parameters) (this *ExistingShamirApproxSS1) {
// 	this = new(ExistingShamirApproxSS1)
// 	this.VanSS = NewVanillaShamirSS(N, T, params)

// 	this.params4cmb, _ = rlwe.NewParameters(params.LogN(), params.Q(), nil, 0, params.HammingWeight(), params.Sigma(), params.RingType(), rlwe.NewScale(1), params.DefaultNTTFlag())
// 	return
// 	//this.params4cmb, _ = rlwe.NewParameters(params.LogN(), params.Q()[:1], nil, 0, params.HammingWeight(), params.Sigma(), params.RingType(), rlwe.NewScale(1), params.DefaultNTTFlag())

// }

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
