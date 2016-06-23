/*
Copyright 2014 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"testing"
	"encoding/base64"
	"fmt"
)

func TestFindLongestLine(t *testing.T){		

	test_in1 := "1\n12\n123\n1234"
	test_in1_decoded, err := base64.StdEncoding.DecodeString(test_in1)
		if err != nil {
		fmt.Println("error:", err)
		return
	}
	test_in2 := "Candy marzipan halvah.\n" + 
				"Tart gummi bears tootsie roll wafer donut tootsie roll. Jujubes gummi bears oat cake toffee cookie tootsie roll.\n" + 
				"Lemon drops soufflé topping oat cake macaroon soufflé ice cream bear claw.\n" + 
				"Dessert jelly carrot cake. Oat cake gummies jelly cupcake icing gummi bears carrot cake carrot cake.\n" +
				"Brownie icing bonbon sweet roll. Marshmallow apple pie chocolate cake toffee. Donut icing biscuit soufflé marshmallow gingerbread.\n" +
				"Soufflé dragée ice cream caramels tiramisu macaroon jelly-o bear claw halvah. Chocolate tart gingerbread biscuit.\n" + 
				"Donut pastry jelly-o.\n"
	test_in2_decoded, err := base64.StdEncoding.DecodeString(test_in2)
		if err != nil {
		fmt.Println("error:", err)
		return
	}

	cases:= []struct { 
		in []byte; want string
	}{ 
		{test_in1_decoded, "1234"},
		{test_in2_decoded, "Brownie icing bonbon sweet roll. Marshmallow apple pie chocolate cake toffee. Donut icing biscuit soufflé marshmallow gingerbread." },
	}

	for _, c := range cases { 
		got := FindLongestLine(c.in)
		if got != c.want { 
				t.Error("FindLongestLine(%q) == %q, want %q", c.in, got, c.want)
		}
	}

}	

