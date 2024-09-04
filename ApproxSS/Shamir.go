package ApproxSS

import (
	"fmt"
	"os"
	"time"

	"github.com/tuneinsight/lattigo/v4/bfv"
	"github.com/tuneinsight/lattigo/v4/drlwe"
	"github.com/tuneinsight/lattigo/v4/ring"
	"github.com/tuneinsight/lattigo/v4/rlwe"
)

type VanillaShamirSS struct {
	//cmb     *MyCombiner
	T int
	N int

	thdizer   *MyThresholdizer
	secret    *rlwe.SecretKey
	secretGen rlwe.KeyGenerator
}

// This function randomly samples a secret and generates necessary parameters for Vanilla Shamir SS
func NewVanillaShamirSS(N, T int, params bfv.Parameters) (vanSS *VanillaShamirSS) {
	vanSS = new(VanillaShamirSS)
	vanSS.T = T
	vanSS.N = N

	vanSS.secretGen = bfv.NewKeyGenerator(params)
	vanSS.secret = vanSS.secretGen.GenSecretKey()
	vanSS.thdizer = NewMyThresholdizer(params.Parameters)

	return
}

func (vanSS *VanillaShamirSS) Share(mes *ring.Poly) (sharesOut []*drlwe.ShamirSecretShare) {
	points := make([]drlwe.ShamirPublicPoint, vanSS.N)
	for i := 0; i < vanSS.N; i++ {
		points[i] = drlwe.ShamirPublicPoint(i + 1)
	}

	if mes != nil {
		secret := vanSS.secretGen.GenSecretKey()
		secret.Value.Q = mes.CopyNew()
		sharesOut = vanSS.thdizer.GenShamirSecretShares(vanSS.T, points, secret)
	} else {
		sharesOut = vanSS.thdizer.GenShamirSecretShares(vanSS.T, points, vanSS.secret)
	}

	return
}

func (vanSS *VanillaShamirSS) ShareThenWrite(mes *ring.Poly, folder string, source *int) (isSuc bool, cleanTime time.Duration, sizeEachShare float64) {

	start := time.Now()
	sharesOut := vanSS.Share(mes)
	cleanTime = time.Since(start)
	for i := 0; i < vanSS.N; i++ {
		var filename string
		if source != nil {
			sourceNumber := *source
			filename = fmt.Sprintf("%s%d%s%d", "mesFrom", sourceNumber, "To", i)
		} else {
			filename = fmt.Sprintf("%s%d", "mesToParty", i)
		}

		skFilename := fmt.Sprintf("%s%s%s%s", "file/", folder, "/", filename)

		shareData, _ := sharesOut[i].MarshalBinary()
		sizeEachShare = float64(len(shareData))
		err := os.WriteFile(skFilename, shareData, 0666)
		if err != nil {
			panic(err)
		}
	}

	return
}

func (vanSS *VanillaShamirSS) Recover() {

}
