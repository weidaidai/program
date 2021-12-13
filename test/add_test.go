package main

import "testing"

func TestAdd(t *testing.T){

	if ans:=add(1,3);ans!=3{

		t.Errorf("1 + 2 expected be 3, but %d got", ans)

	}

}


func TestMul(t *testing.T){

	if ans:=Mul(2,3);ans!=6{

		t.Errorf("2*3 expected be 6, but %d got", ans)

	}

}