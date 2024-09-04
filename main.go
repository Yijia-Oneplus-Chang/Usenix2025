package main

import (
	"ApproxSS/ApproxSS"
	"bufio"
	"flag"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/tuneinsight/lattigo/v4/bfv"
)

func main() {
	// var N = flag.Int("N", 100, "The total number of clients")
	// var T = flag.Int("T", 90, "The threshold")
	var B = flag.Int("B", 16, "The bound of each nonce")

	// paramsLiteral := bfv.ParametersLiteral{
	// 	LogN:     11,
	// 	Q:        []uint64{0x3001}, // 13.5 + 40.4 bits
	// 	Pow2Base: 6,
	// 	T:        0x3001,
	// }

	params, err := bfv.NewParametersFromLiteral(bfv.PN12QP109)
	if err != nil {
		fmt.Println(err)
	}

	flag.Parse()
	//Test our scheme
	// myShamirApproxSS := ApproxSS.NewMyShamirApproxSS(100, 90, params)
	// myShamirApproxSS.VanSS.ShareThenWrite(nil, "skShare", nil)
	// isSucOurScheme, _, _, _, _ := myShamirApproxSS.ApproxRecover(*B)
	// fmt.Println("Our scheme is ", isSucOurScheme)

	//Test existing Shamir scheme with idea 1
	// existingShamir1 := ApproxSS.NewExistingShamirApproxSS1(*N, *T, 1<<paramsLiteral.LogN)
	// existingShamir1.ShareThenWrite(nil)
	// _, isSucExistingShamir1 := existingShamir1.ApproxRecover(*B)
	// fmt.Println("The first idea of Shamir ApproxSS is ", isSucExistingShamir1)

	//Test existing Shamir scheme with idea 2
	// existingShamir2 := ApproxSS.NewExistingShamirApproxSS2(*N, *T, params)
	// existingShamir2.VanSS.ShareThenWrite(nil, "skShare", nil)
	// isSucExistingShamir2, _, _ := existingShamir2.ApproxRecover(*B)
	// fmt.Println("The second idea of Shamir ApproxSS is ", isSucExistingShamir2)

	// repSS := ApproxSS.NewReplicatedSS(*N, *T, params)
	// repSS.ShareThenWrite(nil)
	// isSuc, timeCompRepSS, sizeRepSS := repSS.ApproxRecover(*B)
	// RecordTime("RepSS", timeCompRepSS, sizeRepSS, *N, *T)
	// fmt.Println("The replicated ApproxSS is ", isSuc)

	tRatio := []float64{0.5, 0.7, 0.9}
	for n := 10; n < 110; n = n + 10 {
		for _, tratio := range tRatio {
			t := int(float64(n) * tratio)

			myShamirApproxSS := ApproxSS.NewMyShamirApproxSS(n, t, params)
			myShamirApproxSS.VanSS.ShareThenWrite(nil, "skShare", nil)
			//isSucOurScheme, timeCompOur, timeCompOurNotOnce, sizeOur, sizeOurNotOnce := myShamirApproxSS.ApproxRecover(*B)
			isSucOurScheme, timeCompOur, timeCompOurNotOnce, sizeOur, sizeOurNotOnce := myShamirApproxSS.ApproxRecover4TestTime(*B)
			fmt.Println("Our scheme is ", isSucOurScheme)
			RecordTime("OurOnce", timeCompOur, sizeOur, n, t)
			RecordTime("OurNotOnce", timeCompOurNotOnce, sizeOurNotOnce, n, t)

			existingShamir1 := ApproxSS.NewExistingShamirApproxSS1(n, t, 2048)
			existingShamir1.ShareThenWrite(nil)
			_, isSucExistingShamir1, timeCompShamir1, sizeShamir1 := existingShamir1.ApproxRecover4TestTime(*B)
			fmt.Println("The first idea of Shamir ApproxSS is ", isSucExistingShamir1)
			RecordTime("Shamir1", timeCompShamir1, sizeShamir1, n, t)

			// existingShamir2 := ApproxSS.NewExistingShamirApproxSS2(n, t, params)
			// existingShamir2.VanSS.ShareThenWrite(nil, "skShare", nil)
			// isSucExistingShamir2, timeCompShamir2, sizeShamir2 := existingShamir2.ApproxRecover4TestTime(*B)
			// fmt.Println("The second idea of Shamir ApproxSS is ", isSucExistingShamir2)
			// RecordTime("Shamir2", timeCompShamir2, sizeShamir2, n, t)

			repSS := ApproxSS.NewReplicatedSS(n, t, params)
			repSS.ShareThenWrite4TestTime(nil)
			isSuc, timeCompRepSS, sizeRepSS := repSS.ApproxRecover4TestTime(*B)
			fmt.Println("The replicated ApproxSS is ", isSuc)
			RecordTimeBig("Rep0", timeCompRepSS, sizeRepSS, n, t)

		}

	}

}

func RecordTime(name string, timeComp time.Duration, size float64, N, T int) {

	rate := float64(98)
	timeComm := size / float64(1048576*rate)

	timeTotal := timeComm + timeComp.Seconds()

	var file, file1, file2, file3 *os.File
	var err error

	filename := "TimeTotal"
	if !ApproxSS.CheckFileIsExist(filename) {
		file, err = os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	} else {
		file, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0666)
	}

	if err != nil {
		fmt.Println("文件打开失败", err)
	}

	defer file.Close()

	write := bufio.NewWriter(file)
	mes := fmt.Sprintf("%d%s%d%s%s%s%f\n", N, " ", T, " ", name, " ", timeTotal)
	write.WriteString(mes)

	//Flush将缓存的文件真正写入到文件中
	write.Flush()

	filename1 := "TimeComp"
	if !ApproxSS.CheckFileIsExist(filename1) {
		file1, err = os.OpenFile(filename1, os.O_WRONLY|os.O_CREATE, 0666)
	} else {
		file1, err = os.OpenFile(filename1, os.O_WRONLY|os.O_APPEND, 0666)
	}

	if err != nil {
		fmt.Println("文件打开失败", err)
	}

	defer file1.Close()

	write1 := bufio.NewWriter(file1)
	mes1 := fmt.Sprintf("%d%s%d%s%s%s%f\n", N, " ", T, " ", name, " ", timeComp.Seconds())
	write1.WriteString(mes1)

	//Flush将缓存的文件真正写入到文件中
	write1.Flush()

	filename2 := "TimeComm"
	if !ApproxSS.CheckFileIsExist(filename2) {
		file2, err = os.OpenFile(filename2, os.O_WRONLY|os.O_CREATE, 0666)
	} else {
		file2, err = os.OpenFile(filename2, os.O_WRONLY|os.O_APPEND, 0666)
	}

	if err != nil {
		fmt.Println("文件打开失败", err)
	}

	defer file2.Close()

	write2 := bufio.NewWriter(file2)
	mes2 := fmt.Sprintf("%d%s%d%s%s%s%f\n", N, " ", T, " ", name, " ", timeComm)
	write2.WriteString(mes2)

	//Flush将缓存的文件真正写入到文件中
	write2.Flush()

	filename3 := "SizeComm"
	if !ApproxSS.CheckFileIsExist(filename3) {
		file3, err = os.OpenFile(filename3, os.O_WRONLY|os.O_CREATE, 0666)
	} else {
		file3, err = os.OpenFile(filename3, os.O_WRONLY|os.O_APPEND, 0666)
	}

	if err != nil {
		fmt.Println("文件打开失败", err)
	}

	defer file3.Close()

	write3 := bufio.NewWriter(file3)
	mes3 := fmt.Sprintf("%d%s%d%s%s%s%f\n", N, " ", T, " ", name, " ", size)
	write3.WriteString(mes3)

	//Flush将缓存的文件真正写入到文件中
	write3.Flush()

}

func RecordTimeBig(name string, timeComp, size *big.Float, N, T int) {

	rate := float64(98)
	timeComm := big.NewFloat(0)
	timeComm.Quo(size, new(big.Float).SetFloat64(float64(1048576*rate)))

	timeTotal := new(big.Float).Add(timeComm, timeComp)

	var file, file1, file2, file3 *os.File
	var err error

	filename := "TimeTotal"
	if !ApproxSS.CheckFileIsExist(filename) {
		file, err = os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	} else {
		file, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0666)
	}

	if err != nil {
		fmt.Println("文件打开失败", err)
	}

	defer file.Close()

	write := bufio.NewWriter(file)
	mes := fmt.Sprintf("%d%s%d%s%s%s%s\n", N, " ", T, " ", name, " ", timeTotal.String())
	write.WriteString(mes)

	//Flush将缓存的文件真正写入到文件中
	write.Flush()

	filename1 := "TimeComp"
	if !ApproxSS.CheckFileIsExist(filename1) {
		file1, err = os.OpenFile(filename1, os.O_WRONLY|os.O_CREATE, 0666)
	} else {
		file1, err = os.OpenFile(filename1, os.O_WRONLY|os.O_APPEND, 0666)
	}

	if err != nil {
		fmt.Println("文件打开失败", err)
	}

	defer file1.Close()

	write1 := bufio.NewWriter(file1)
	mes1 := fmt.Sprintf("%d%s%d%s%s%s%s\n", N, " ", T, " ", name, " ", timeComp.String())
	write1.WriteString(mes1)

	//Flush将缓存的文件真正写入到文件中
	write1.Flush()

	filename2 := "TimeComm"
	if !ApproxSS.CheckFileIsExist(filename2) {
		file2, err = os.OpenFile(filename2, os.O_WRONLY|os.O_CREATE, 0666)
	} else {
		file2, err = os.OpenFile(filename2, os.O_WRONLY|os.O_APPEND, 0666)
	}

	if err != nil {
		fmt.Println("文件打开失败", err)
	}

	defer file2.Close()

	write2 := bufio.NewWriter(file2)
	mes2 := fmt.Sprintf("%d%s%d%s%s%s%s\n", N, " ", T, " ", name, " ", timeComm.String())
	write2.WriteString(mes2)

	//Flush将缓存的文件真正写入到文件中
	write2.Flush()

	filename3 := "SizeComm"
	if !ApproxSS.CheckFileIsExist(filename3) {
		file3, err = os.OpenFile(filename3, os.O_WRONLY|os.O_CREATE, 0666)
	} else {
		file3, err = os.OpenFile(filename3, os.O_WRONLY|os.O_APPEND, 0666)
	}

	if err != nil {
		fmt.Println("文件打开失败", err)
	}

	defer file3.Close()

	write3 := bufio.NewWriter(file3)
	mes3 := fmt.Sprintf("%d%s%d%s%s%s%s\n", N, " ", T, " ", name, " ", size.String())
	write3.WriteString(mes3)

	//Flush将缓存的文件真正写入到文件中
	write3.Flush()

}
