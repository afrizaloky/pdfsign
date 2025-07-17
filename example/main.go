package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/afrizaloky/pdfsign/verify"
)

func VerifyPDF(ctx context.Context, pdfData []byte) (*verify.Response, error) {

	pdfSize := int64(len(pdfData))
	reader := bytes.NewReader(pdfData)
	resp, err := verify.Reader(reader, pdfSize)
	if err != nil {
		return resp, err
	}

	if resp.Error != "" {
		return resp, errors.New(resp.Error)
	}
	return resp, nil
}

func main() {
	pdf, err := os.ReadFile("/home/afrizal/Documents/SharedFolder/ecdsa.pdf")
	if err != nil {
		log.Fatal(err)
	}
	r, err := VerifyPDF(context.Background(), pdf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r.Error)
	fmt.Println("hehe")
}
