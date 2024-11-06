# Console bar

## What is the "Console Bar" 

A pet-project :B

## How could you use the project?

1. Create a derectory (folder)

2. go to the folder in the console 
```shell
cd path_to_your_folder
```

3. write in the console 
```shell
go mod init your_project_name
```

4. install the project 

```shell 
go get github.com/VandiKond/ConsoleBar
```

## Standart usage 

```golang
package main

import coffe_construcurs "github.com/VandiKond/ConsoleBar/constructurs/coffe"

func main() {
	coffe_construcurs.NewLatte()
}
```

- It could change in different versions

## Interfaces 

### Additive 

- It's an interface for additives for drinks (example: sugar, cinnamon)

- Methods

1. GetGrams 
    - Gets the amount of the additive (in grams)
    - Example : 
    ```golang
    func (a Additive) GetGrams() float64 {
        return a.Grams
    }
    ```
    - structure for this example should have parameter `Grams`

2. GetType
    - Gets the type of the additive
    - Example : 
    ```golang
    func (a Additive) GetType() string {
	    return "additive"
    }
    ```
    - replace "additive" to your structure name

3. AddMore
    - Ads more grams to the additive
    - Example :
    ```golang 
    func (a Additive) AddMore(grams float64) Additive {
        a.Grams += grams
        return a
    }
    ```
    - structure for this example should have parameter `Grams`

### Coffee 

- It's an interface for coffee (example: latte, cappuchino, coffee with milk)

- Methods 

1. AddSyrop
    - Ads syrop to the coffee
    - Example:
    ```golang
    func (c Coffee) AddSyrop(syrop syrops.Syrop) (Coffee, error) {
        // Checks are syrops similar
        if c.Syrop.GetType() != syrop.GetType() {
            return l, Errors.NewError(coffeerrors.TDS, fmt.Sprintf("%s is not %s", c.Syrop.GetType(), syrop.GetType()))
        }
        // Ads syrop
        c.Syrop = l.Syrop.AddMore(syrop.GetMl())

        // returns the coffee
        return l, nil
    }
    ```
    - Expects that the sturcture you are using has a parametor `Syrop` 

2. AddAdditive
    - Ads an additive to the coffee
    - Example:
    ```golang
    func (c Coffee) AddAdditive(additive additives.Additive) (Coffee, error) {
        // Ads a new element to the slice of additives
        c.Additives = append(c.Additives, additive)

        // returns the coffee
        return c, nil
    }
    ```
    - Expects that the sturcture you are using has a parametor `Additives` *(slice of Addective)*

### Syrops

- It's an interface for syrops that you can add to drinks (example: vanila syrop, choclate syrop, salted caramel syrup)

- Methods 

1. GetMl 
    - Gets the amount of the syrop (in ml)
    - Example : 
    ```golang
    func (s Syrop) GetMl() float64 {
        return s.Ml
    }
    ```
    - structure for this example should have parameter `Ml`

2. GetType
    - Gets the type of the syrop
    - Example : 
    ```golang
    func (s Syrop) GetType() string {
	    return "syrop"
    }
    ```
    - replace "syrop" to your structure name

3. AddMore
    - Ads more grams to the additive
    - Example :
    ```golang 
    func (s Syrop) AddMore(grams float64) Syrop {
        v.Ml += ml
        return v
    }
    ```
    - structure for this example should have parameter `Ml`

## Error types

### Coffe errors

- TDS : cant add two different syrops (two different syrops)
    - indicates that you can't mix two sirops 

- NEC : not enougth coffe
    - indicates that you can't have zero or negative amount of coffee

### Syrop errors 

- UST : unknown syrop type
    - indicates, that the program dosen't know the syrop type

## Other 

- actualy there are other more tyny methods that i haven't managed

- In some time i would talk about the created functionality :B