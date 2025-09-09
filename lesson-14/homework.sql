SELECT model, speed, hd
FROM PC
WHERE price < 500;

SELECT DISTINCT(maker)
FROM Product
WHERE type = 'Printer';

SELECT model, ram, screen FROM Laptop where price > 1000;

SELECT * FROM Printer WHERE color = 'y';

SELECT model, speed, hd FROM PC where price < 600 AND (cd = '12x' OR cd = '24x');

SELECT DISTINCT p.maker, l.speed
FROM Product p
         JOIN Laptop l ON p.model = l.model
WHERE l.hd >= 10;


SELECT p.model, pc.price
FROM Product p
         JOIN PC pc ON p.model = pc.model
WHERE p.maker = 'B'
UNION
SELECT p.model, l.price
FROM Product p
         JOIN Laptop l ON p.model = l.model
WHERE p.maker = 'B'
UNION
SELECT p.model, pr.price
FROM Product p
         JOIN Printer pr ON p.model = pr.model
WHERE p.maker = 'B';

SELECT DISTINCT maker
FROM Product
WHERE type = 'PC'
  AND maker NOT IN (
    SELECT maker
    FROM Product
    WHERE type = 'Laptop'
);

SELECT DISTINCT p.maker
FROM Product p
         JOIN PC pc ON p.model = pc.model
WHERE pc.speed >= 450;

SELECT model, price
FROM Printer
WHERE price = (SELECT MAX(price) FROM Printer);
