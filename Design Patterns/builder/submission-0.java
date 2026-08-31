class Meal {

    private double cost;
    private boolean takeOut;
    private String main;
    private String drink;

    double getCost() {
        return this.cost;
    }

    boolean getTakeOut() {
        return this.takeOut;
    }

    String getMain() {
        return this.main;
    }

    String getDrink() {
        return this.drink;
    }

    void setCost(double cost) {
        this.cost = cost;
    }

    void setTakeOut(boolean takeOut) {
        this.takeOut = takeOut;
    }

    void setMain(String main) {
        this.main = main;
    }

    void setDrink(String drink) {
        this.drink = drink;
    }
}

class MealBuilder {
    private double cost;
    private boolean takeOut;
    private String main;
    private String drink;

    public MealBuilder() {}

    public MealBuilder addCost(double cost) {
        this.cost = cost;
        return this;
    }

    public MealBuilder addTakeOut(boolean takeOut) {
        this.takeOut = takeOut;
        return this;
    }

    public MealBuilder addMainCourse(String main) {
        this.main = main;
        return this;
    }

    public MealBuilder addDrink(String drink) {
        this.drink = drink;
        return this;
    }

    Meal build() {
        Meal meal = new Meal();
        meal.setCost(this.cost);
        meal.setTakeOut(this.takeOut);
        meal.setMain(this.main);
        meal.setDrink(this.drink);

        return meal;
    }
}
