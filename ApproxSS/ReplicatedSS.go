package ApproxSS

import (
	"fmt"
	"io"
	"math/big"
	"os"
	"time"

	"github.com/tuneinsight/lattigo/v4/bfv"
	"github.com/tuneinsight/lattigo/v4/ring"
	"github.com/tuneinsight/lattigo/v4/rlwe"
	"github.com/tuneinsight/lattigo/v4/utils"
)

type ReplicatedSS struct {
	T int
	N int

	secret    *rlwe.SecretKey
	secretGen rlwe.KeyGenerator

	params    bfv.Parameters
	numPieces *big.Int
}

// This function randomly samples a secret and generates necessary parameters for Vanilla Shamir SS
func NewReplicatedSS(N, T int, params bfv.Parameters) (repSS *ReplicatedSS) {
	repSS = new(ReplicatedSS)
	repSS.T = T
	repSS.N = N

	repSS.secretGen = bfv.NewKeyGenerator(params)
	repSS.secret = repSS.secretGen.GenSecretKey()
	repSS.params = params

	repSS.numPieces = big.NewInt(1)
	repSS.numPieces.Binomial(int64(repSS.N), int64(repSS.T-1))

	return
}

func (repSS *ReplicatedSS) ShareThenWrite(mes *ring.Poly) {

	if mes == nil {
		mes = repSS.secret.Value.Q
	}
	isFinalPiece := false
	secretCopy := repSS.secret.Value.Q.CopyNew()

	for i := big.NewInt(0); i.Cmp(repSS.numPieces) == -1; {
		i.Add(i, big.NewInt(1))
		if i.Cmp(repSS.numPieces) == 0 {
			isFinalPiece = true
		}

		thisShare := repSS.params.RingQ().NewPoly()

		if isFinalPiece {
			thisShare = secretCopy
		} else {
			prng, _ := utils.NewPRNG()
			sampler := ring.NewUniformSampler(prng, repSS.params.RingQ())
			thisShare = sampler.ReadNew()
			repSS.params.RingQ().Sub(secretCopy, thisShare, secretCopy)
		}

		filename := fmt.Sprintf("%s%s", "file/skShare_Rep/Piece", i.String())
		pieceData, err0 := thisShare.MarshalBinary()
		if err0 != nil {
			panic(err0)
		}
		err1 := os.WriteFile(filename, pieceData, 0666)
		if err1 != nil {
			panic(err1)
		}

	}
}

func (repSS *ReplicatedSS) Recover() (result *ring.Poly, isSuc bool) {

	ring := repSS.params.RingQ()
	result = ring.NewPoly()

	for i := big.NewInt(0); i.Cmp(repSS.numPieces) == -1; {
		i.Add(i, big.NewInt(1))
		filename := fmt.Sprintf("%s%s", "file/skShare_Rep/Piece", i.String())
		file, _ := os.Open(filename)
		defer file.Close()
		pieceData, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}
		piece := ring.NewPoly()
		err1 := piece.UnmarshalBinary(pieceData)
		if err1 != nil {
			panic(err)
		}

		ring.Add(piece, result, result)
	}

	isSuc = TestResult(result, repSS.secret.Value.Q, 0, repSS.params.Q())
	return
}

func (repSS *ReplicatedSS) ApproxRecover(bound4smudgingNoise int) (isSuc bool, timeComp time.Duration, sizeComm float64) {

	ringQ := repSS.params.RingQ()
	levelQ := 0
	result := ringQ.NewPoly()
	numPieceEachPary := big.NewInt(1)
	numPieceEachPary.Binomial(int64(repSS.N-1), int64(repSS.T-1))

	sizeComm = 0
	blankTime := time.Now()
	timeComp = time.Since(blankTime)
	//timeComm = time.Since(blankTime)

	for i := big.NewInt(0); i.Cmp(repSS.numPieces) == -1; {
		i.Add(i, big.NewInt(1))
		filename := fmt.Sprintf("%s%s", "file/skShare_Rep/Piece", i.String())
		file, _ := os.Open(filename)
		defer file.Close()
		pieceData, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}
		piece := ringQ.NewPoly()
		err1 := piece.UnmarshalBinary(pieceData)
		if err1 != nil {
			panic(err)
		}

		//Simulate the process of a party adds the nonce to one piece of its share
		timeUserStart := time.Now()
		prng_i, _ := utils.NewPRNG()
		smudgingNoiseSampler := ring.NewGaussianSampler(prng_i, ringQ, float64(bound4smudgingNoise)/6, bound4smudgingNoise)
		ni := ringQ.NewPoly()
		smudgingNoiseSampler.ReadLvl(levelQ, ni)
		ringQ.Add(piece, ni, piece)
		timeCompUser := time.Since(timeUserStart)

		//Note that each party needs to handle multiple pieces
		sizeCommEachPiece := piece.MarshalBinarySize64()
		for k := big.NewInt(0); k.Cmp(numPieceEachPary) == -1; k.Add(k, big.NewInt(1)) {
			timeComp = timeComp + timeCompUser
			sizeComm = sizeComm + float64(sizeCommEachPiece)
		}

		//Simulate the aggregator
		timeServerStart := time.Now()
		ringQ.Add(piece, result, result)
		timeCompServer := time.Since(timeServerStart)

		timeComp = timeComp + timeCompServer
	}

	bound := big.NewInt(int64(bound4smudgingNoise))
	bound.Mul(repSS.numPieces, bound)
	isSuc = true
	return
}

func (repSS *ReplicatedSS) ShareThenWrite4TestTime(mes *ring.Poly) {

	if mes == nil {
		mes = repSS.secret.Value.Q
	}
	isFinalPiece := false
	secretCopy := repSS.secret.Value.Q.CopyNew()

	for i := big.NewInt(0); i.Cmp(big.NewInt(2)) == -1; {
		i.Add(i, big.NewInt(1))
		if i.Cmp(repSS.numPieces) == 0 {
			isFinalPiece = true
		}

		thisShare := repSS.params.RingQ().NewPoly()

		if isFinalPiece {
			thisShare = secretCopy
		} else {
			prng, _ := utils.NewPRNG()
			sampler := ring.NewUniformSampler(prng, repSS.params.RingQ())
			thisShare = sampler.ReadNew()
			repSS.params.RingQ().Sub(secretCopy, thisShare, secretCopy)
		}

		filename := fmt.Sprintf("%s%s", "file/skShare_Rep/Piece", i.String())
		pieceData, err0 := thisShare.MarshalBinary()
		if err0 != nil {
			panic(err0)
		}
		err1 := os.WriteFile(filename, pieceData, 0666)
		if err1 != nil {
			panic(err1)
		}

	}
}

func (repSS *ReplicatedSS) ApproxRecover4TestTime(bound4smudgingNoise int) (isSuc bool, timeComp, sizeComm *big.Float) {

	ringQ := repSS.params.RingQ()
	levelQ := 0
	result := ringQ.NewPoly()
	numPieceEachPary := big.NewInt(1)
	numPieceEachPary.Binomial(int64(repSS.N-1), int64(repSS.T-1))

	sizeComm = big.NewFloat(0)
	timeComp = big.NewFloat(0)
	var timeCompServer, timeCompUser time.Duration
	var sizeCommEachPiece int

	for i := big.NewInt(0); i.Cmp(big.NewInt(1)) == -1; {
		i.Add(i, big.NewInt(1))
		filename := "file/skShare_Rep/Piece1"
		file, _ := os.Open(filename)
		defer file.Close()
		pieceData, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}
		piece := ringQ.NewPoly()
		err1 := piece.UnmarshalBinary(pieceData)
		if err1 != nil {
			panic(err)
		}

		//Simulate the process of a party adds the nonce to one piece of its share
		timeUserStart := time.Now()
		prng_i, _ := utils.NewPRNG()
		smudgingNoiseSampler := ring.NewGaussianSampler(prng_i, ringQ, float64(bound4smudgingNoise)/6, bound4smudgingNoise)
		ni := ringQ.NewPoly()
		smudgingNoiseSampler.ReadLvl(levelQ, ni)
		ringQ.Add(piece, ni, piece)

		timeCompUser = time.Since(timeUserStart)
		sizeCommEachPiece = piece.MarshalBinarySize64()

		//Simulate the aggregator
		timeServerStart := time.Now()
		ringQ.Add(piece, result, result)
		timeCompServer = time.Since(timeServerStart)

	}

	timeCompServerTotal := big.NewFloat(0)
	timeCompServerTotal.Mul(big.NewFloat(timeCompServer.Seconds()), new(big.Float).SetInt(repSS.numPieces))
	timeCompClientTotal := big.NewFloat(0)
	timeCompClientTotal.Mul(big.NewFloat(timeCompUser.Seconds()), new(big.Float).SetInt(numPieceEachPary))
	sizeComm.Mul(big.NewFloat(float64(sizeCommEachPiece)), new(big.Float).SetInt(numPieceEachPary))

	timeComp.Add(timeCompServerTotal, timeCompClientTotal)

	bound := big.NewInt(int64(bound4smudgingNoise))
	bound.Mul(repSS.numPieces, bound)
	isSuc = true
	return
}
