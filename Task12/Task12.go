package main

import (
	"fmt"
	"strings"
)

type Ingredient struct {
	Name   string
	Amount string 
}

type Recipe struct {
	ID          int
	Title       string      
	Description string        
	Ingredients []Ingredient  
	Steps       []string      
	Category    string         
	PrepTime    int           
}

type RecipeBook struct {
	recipes       []Recipe
	nextRecipeID  int
}

func (rb *RecipeBook) AddRecipe(title, description string, ingredients []Ingredient, steps []string, category string, prepTime int) {
	newRecipe := Recipe{
		ID:          rb.nextRecipeID,
		Title:       title,
		Description: description,
		Ingredients: ingredients,
		Steps:       steps,
		Category:    category,
		PrepTime:    prepTime,
	}
	rb.recipes = append(rb.recipes, newRecipe)
	rb.nextRecipeID++
	fmt.Println("Добавлен рецепт:", title)
}

func (rb *RecipeBook) PrintAllRecipes() {
	for _, recipe := range rb.recipes {
		fmt.Println("--------------------------------------------------")
		fmt.Println("ID:", recipe.ID)
		fmt.Println("Название:", recipe.Title)
		fmt.Println("Описание:", recipe.Description)
		fmt.Println("Категория:", recipe.Category)
		fmt.Println("Время приготовления:", recipe.PrepTime, "мин")
		fmt.Println("Ингредиенты:")
		for _, ing := range recipe.Ingredients {
			fmt.Println("-", ing.Name, ":", ing.Amount)
		}
		fmt.Println("Этапы приготовления:")
		for i, step := range recipe.Steps {
			fmt.Println(i+1, "-", step)
		}
	}
	fmt.Println("--------------------------------------------------")
}

func (rb *RecipeBook) SearchByIngredient() {
	fmt.Println("Введите название ингредиента для поиска:")
	var input string
	fmt.Scanln(&input)
	input = strings.ToLower(input)
	found := false
	for _, recipe := range rb.recipes {
		for _, ing := range recipe.Ingredients {
			if strings.Contains(strings.ToLower(ing.Name), input) {
				fmt.Println("Найден рецепт:", recipe.Title)
				fmt.Println("Категория:", recipe.Category)
				fmt.Println("Время приготовления:", recipe.PrepTime, "мин")
				fmt.Println("--------------------------------------------------")
				found = true
				break
			}
		}
	}

	if !found {
		fmt.Println("Рецепты с таким ингредиентом не найдены.")
	}
}

func (rb *RecipeBook) FilterByCategory() {
	fmt.Println("Введите категорию для фильтрации:")
	var input string
	fmt.Scanln(&input)
	input = strings.ToLower(input)
	found := false

	for _, recipe := range rb.recipes {
		if strings.ToLower(recipe.Category) == input {
			fmt.Println("Рецепт:", recipe.Title)
			fmt.Println("Описание:", recipe.Description)
			fmt.Println("Время приготовления:", recipe.PrepTime, "мин")
			fmt.Println("--------------------------------------------------")
			found = true
		}
	}

	if !found {
		fmt.Println("Рецепты с такой категорией не найдены.")
	}
}

func (rb *RecipeBook) FindLongestRecipe() {
	if len(rb.recipes) == 0 {
		fmt.Println("В книге нет рецептов.")
		return
	}

	longest := rb.recipes[0]
	for _, recipe := range rb.recipes {
		if recipe.PrepTime > longest.PrepTime {
			longest = recipe
		}
	}

	fmt.Println("Самый долгий рецепт:")
	fmt.Println("Название:", longest.Title)
	fmt.Println("Время приготовления:", longest.PrepTime, "мин")
	fmt.Println("Категория:", longest.Category)
	fmt.Println("--------------------------------------------------")
}

func main() {
	rb := RecipeBook{
		recipes:      []Recipe{},
		nextRecipeID: 1,
	}
	rb.AddRecipe(
		"Блины",
		"Классические русские блины",
		[]Ingredient{
			{"Молоко", "500 мл"},
			{"Мука", "250 г"},
			{"Яйца", "2 шт"},
			{"Сахар", "2 ст.л."},
		},
		[]string{
			"Смешать яйца с сахаром",
			"Добавить молоко и муку, перемешать",
			"Жарить на сковороде до готовности",
		},
		"завтрак",
		30,
	)
	rb.AddRecipe(
		"Шоколадный торт",
		"Вкусный десерт для всей семьи",
		[]Ingredient{
			{"Шоколад", "200 г"},
			{"Мука", "300 г"},
			{"Яйца", "3 шт"},
			{"Сахар", "150 г"},
		},
		[]string{
			"Растопить шоколад",
			"Смешать все ингредиенты",
			"Выпекать 40 минут при 180°C",
		},
		"десерт",
		60,
	)
	rb.PrintAllRecipes()
	rb.SearchByIngredient()
	rb.FilterByCategory()
	rb.FindLongestRecipe()
}
