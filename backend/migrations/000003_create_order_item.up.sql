CREATE TABLE IF NOT EXISTS ORDER_ITEM(
   id serial PRIMARY KEY,
   id_pizza INT NOT NULL,
   quantity INT NOT NULL,
   id_order INT NOT NULL,
      CONSTRAINT fk_pizza
        FOREIGN KEY(id_pizza) 
	        REFERENCES PIZZA(id),
      CONSTRAINT fk_order
        FOREIGN KEY(id_order) 
	        REFERENCES ORDER_PIZZA(id)
);