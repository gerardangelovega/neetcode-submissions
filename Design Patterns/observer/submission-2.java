interface Observer {
    void notify(String itemName);
}

class Customer implements Observer {
    private String name;
    private int notifications;

    public Customer(String name) {
        this.name = name;
        this.notifications = 0;
    }

    public void notify(String itemName) {
        this.notifications += 1;
    }

    public int countNotifications() {
        return this.notifications;
    }
}

class OnlineStoreItem {
    private String itemName;
    private int stock;
    private List<Observer> observers = new ArrayList<Observer>();

    public OnlineStoreItem(String itemName, int stock) {
        this.itemName = itemName;
        this.stock = stock;
    }

    public void subscribe(Observer observer) {
        this.observers.add(observer);
    }

    public void unsubscribe(Observer observer) {
        for (int i = 0; i < this.observers.size(); i++) {
            if (this.observers.get(i) == observer) {
                this.observers.remove(i);
                break;
            }
        }
    }

    public void updateStock(int newStock) {
        if (this.stock != 0) {
            this.stock = newStock;
            return;
        }
        this.stock = newStock;
        for (Observer observer : this.observers) {
            observer.notify(this.itemName);
        }
    }
}
