package personChannel

import(
	"github.com/kd2718/gomarshal/person"
)

func PersonChannel(p *person.Person) <- chan person.Person {
	//var out chan interface{}
	guy := *p
	out := make(chan person.Person)

	go func() {
		defer close(out)
		for ; guy.Age < person.Age(150); guy.Birthday(){
			out <- guy
		}
    }()

	return out

}

