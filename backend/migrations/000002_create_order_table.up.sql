CREATE TABLE IF NOT EXISTS ORDER_PIZZA(
   id serial PRIMARY KEY,
   id_pizza INT NOT NULL,
   quantity INT NOT NULL,
      CONSTRAINT fk_pizza
        FOREIGN KEY(id_pizza) 
	        REFERENCES PIZZA(id)
);