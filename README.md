# Console bar

## What is the "Console Bar" 

A pet-project :B

## How could you use the project?

1. Create a directory (folder)

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

## Standard usage 

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
    - Adds more grams to the additive
    - Example :
    ```golang 
    func (a Additive) AddMore(grams float64) Additive {
        a.Grams += grams
        return a
    }
    ```
    - structure for this example should have parameter `Grams`

### Coffee 

- It's an interface for coffee (example: latte, cappuccino, coffee with milk)

- Methods 

1. Addsyrup
    - Adds syrup to the coffee
    - Example:
    ```golang
    func (c Coffee) AddSyrup(syrup syrups.Syrup) (Coffee, error) {
        // Checks are syrups similar
        if c.Syrup.GetType() != syrup.GetType() {
            return l, Errors.NewError(coffeerrors.TDS, fmt.Sprintf("%s is not %s", c.Syrup.GetType(), syrup.GetType()))
        }
        // Adds syrup
        c.Syrup = l.Syrup.AddMore(syrup.GetMl())

        // returns the coffee
        return l, nil
    }
    ```
    - Expects that the structure you are using has a parameter `Syrup` 

2. AddAdditive
    - Adds an additive to the coffee
    - Example:
    ```golang
    func (c Coffee) AddAdditive(additive additives.Additive) (Coffee, error) {
        // Adds a new element to the slice of additives
        c.Additives = append(c.Additives, additive)

        // returns the coffee
        return c, nil
    }
    ```
    - Expects that the structure you are using has a parameter `Additives` *(slice of Addective)*

### Syrups

- It's an interface for syrups that you can add to drinks (example: vanilla syrup, chocolate syrup, salted caramel syrup)

- Methods 

1. GetMl 
    - Gets the amount of the syrup (in ml)
    - Example : 
    ```golang
    func (s Syrup) GetMl() float64 {
        return s.Ml
    }
    ```
    - structure for this example should have parameter `Ml`

2. GetType
    - Gets the type of the syrup
    - Example : 
    ```golang
    func (s Syrup) GetType() string {
	    return "syrup"
    }
    ```
    - replace "syrup" to your structure name

3. AddMore
    - Adds more grams to the additive
    - Example :
    ```golang 
    func (s Syrup) AddMore(grams float64) Syrup {
        v.Ml += ml
        return v
    }
    ```
    - structure for this example should have parameter `Ml`

## Error types

### Coffee errors

- TDS : cant add two different syrups (two different syrups)
    - indicates that you can't mix two sirops 

- NEC : not enough coffee
    - indicates that you can't have zero or negative amount of coffee

### Syrup errors 

- UST : unknown syrup type
    - indicates, that the program doesn't know the syrup type

## Other 

- actually there are other more methods that i haven't managed

- In some time i would talk about the created functionality :B