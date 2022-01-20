package main

import (
	"bytes"
	"fmt"
	"github.com/v2fly/VSign/insmgr"
	"github.com/v2fly/VSign/instimp"
	"github.com/v2fly/VSign/sign"
	"github.com/v2fly/VSign/signerVerify"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	argoffset := 1

	outins := insmgr.NewOutputInsMgr(os.Stdout)
	switch os.Args[0+argoffset] {
	case "gen":
		switch os.Args[1+argoffset] {
		case "sort":
			f, _ := os.Open("sigb")
			insmgr.SortAll(f, os.Stdout)
		case "version":
			insmgr.NewYieldSingle(instimp.NewVersionIns(os.Args[2+argoffset])).InstructionYield(outins)
		case "project":
			insmgr.NewYieldSingle(instimp.NewProjectIns(os.Args[2+argoffset])).InstructionYield(outins)
		case "file":
			instimp.NewFileBasedInsYield(os.Args[2+argoffset]).InstructionYield(outins)
		case "key":
			prv, pub := sign.GenerateKeyFromSeed(os.Args[2+argoffset], os.Args[3+argoffset])
			ioutil.WriteFile("prv.sec", prv, 0600)
			ioutil.WriteFile("pub.pub", pub, 0600)
		}
		return
	case "sign":
		w, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		ins := insmgr.ReadAllIns(bytes.NewReader(w))
		// Check
		_ = ins
		key, err := ioutil.ReadFile("prv.sec")
		if err != nil {
			panic(err)
		}
		password := os.Args[1+argoffset]

		sw, err := sign.Sign(key, password, w)
		if err != nil {
			panic(err)
		}
		io.Copy(os.Stdout, bytes.NewReader(sw))
		fmt.Println()
		return
	case "verify":
		switch os.Args[1+argoffset] {
		case "skip":
		default:
		}
		return
	case "check":
		sigfile := os.Args[1+argoffset]
		fileslist := os.Args[2+argoffset:]
		file, err := os.Open(sigfile)
		if err != nil {
			fmt.Printf("Unable to open signature file %v\n", file)
			os.Exit(-1)
		}
		err = signerVerify.OutputAndJudge(signerVerify.CheckSignaturesV2Fly(file, fileslist))
		file.Close()
		if err != nil {
			os.Exit(-2)
		}
		os.Exit(0)
		return

	}

}
