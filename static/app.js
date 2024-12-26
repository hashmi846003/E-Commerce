document.addEventListener("DOMContentLoaded", () => {
    const productsList = document.getElementById("products-list");
    const addProductForm = document.getElementById("add-product-form");
    const cartList = document.getElementById("cart-list");
  
    const API_BASE = "http://localhost:8080";
  
    // Fetch and display all products
    async function fetchProducts() {
      try {
        const response = await fetch(`${API_BASE}/products`);
        const products = await response.json();
        productsList.innerHTML = "";
        products.forEach(product => {
          const li = document.createElement("li");
          li.textContent = `${product.name} - $${product.price} (Stock: ${product.stock})`;
          productsList.appendChild(li);
        });
      } catch (err) {
        console.error("Error fetching products:", err);
      }
    }
  
    // Add a new product
    addProductForm.addEventListener("submit", async (e) => {
      e.preventDefault();
      const name = document.getElementById("product-name").value;
      const price = parseFloat(document.getElementById("product-price").value);
      const stock = parseInt(document.getElementById("product-stock").value);
  
      try {
        const response = await fetch(`${API_BASE}/products`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ name, price, stock }),
        });
        if (response.ok) {
          alert("Product added successfully!");
          fetchProducts();
          addProductForm.reset();
        } else {
          alert("Failed to add product.");
        }
      } catch (err) {
        console.error("Error adding product:", err);
      }
    });
  
    // Fetch and display the cart (example)
    async function fetchCart() {
      try {
        const response = await fetch(`${API_BASE}/cart`);
        const cart = await response.json();
        cartList.innerHTML = "";
        cart.forEach(item => {
          const li = document.createElement("li");
          li.textContent = `${item.product_name} - Quantity: ${item.quantity}`;
          cartList.appendChild(li);
        });
      } catch (err) {
        console.error("Error fetching cart:", err);
      }
    }
  
    // Initialize
    fetchProducts();
    fetchCart();
  });
  