import { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";

export default function App() {
  const [health, setHealth] = useState("Checking backend...");
  const [coffees, setCoffees] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  // form data state for creating coffee
  const [formData, setFormData] = useState({
    name: "",
    roast: "default medium",
    image: "",
    region: "",
    price: "",
    grind_unit: 250,
  });

  useEffect(() => {
    async function loadData() {
      try {
        const [healthRes, coffeesRes] = await Promise.all([
          fetch("/health"),
          fetch("/api/coffees/all"),
        ]);

        const healthText = await healthRes.text();
        setHealth(healthText || "Backend reachable.");

        if (!coffeesRes.ok) {
          throw new Error("Failed to load coffees from API.");
        }

        const coffeesData = await coffeesRes.json();
        setCoffees(coffeesData?.coffees ?? []);
      } catch (err) {
        setError(err instanceof Error ? err.message : "Unexpected error");
      } finally {
        setLoading(false);
      }
    }

    loadData();
  }, []);

  // create coffee
  async function handleCreate(coffeeInput) {
    setError("");
    try {
      const res = await fetch("/api/coffees/create", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(coffeeInput),
      });

      if (!res.ok) {
        throw new Error("Failed to create coffee.");
      }

      const data = await res.json();
      const createdCoffee = data?.coffee ?? coffeeInput;

      setCoffees((prev) => [createdCoffee, ...prev]);

      // reset form
      setFormData({
        name: "",
        roast: "Medium",
        image: "",
        region: "",
        price: "",
        grind_unit: 250,
      });
    } catch (err) {
      setError(err instanceof Error ? err.message : "Create failed");
    }
  }

  // handle form input change
  function handleFormChange(e) {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]:
        name === "price" || name === "grind_unit" ? parseFloat(value) : value,
    }));
  }

  // handle form submit
  function handleFormSubmit(e) {
    e.preventDefault();
    if (!formData.name || !formData.region) {
      setError("Name and region are required");
      return;
    }
    handleCreate(formData);
  }

  // delete coffee by id (preferred) or by name fallback
  async function handleDelete(coffee) {
    setError("");

    const hasID = Boolean(coffee?.id);
    const encodedId = hasID ? encodeURIComponent(coffee.id) : "";
    const encodedName = !hasID ? encodeURIComponent(coffee?.name ?? "") : "";
    const endPoint = encodedId
      ? `/api/coffees/del/id/${encodedId}`
      : `/api/coffees/del/name/${encodedName}`;

    try {
      const res = await fetch(endPoint, { method: "DELETE" });
      if (!res.ok) {
        throw new Error("failed to delete coffee ");
      }
      // removing that coffee from existence
      setCoffees((prevArr) => {
        return prevArr.filter((item) => {
          if (hasID) {
            return item.id !== coffee.id;
          }
          return item.name !== coffee.name;
        });
      });
    } catch (err) {
      setError(err instanceof Error ? err.message : "Delete failed");
    }
  }

  return (
    <div className="page">
      <header className="hero">
        <h1>Coffee Hub</h1>
        <nav className="hero-navbar">
          <ul className="hero-navbar-ul">
            <li className="hero-navbar-ul-li">
              <Link to="/home">Home</Link>
            </li>
            <li className="hero-navbar-ul-li">
              <Link to="/collections">Collections</Link>
            </li>
            <li className="hero-navbar-ul-li">
              <Link to="/home">Shop</Link>
            </li>
            <li className="hero-navbar-ul-li">
              <Link to="/home">Help</Link>
            </li>
          </ul>
        </nav>
        {/* <p>{health}l</p> */}
      </header>

      <main className="content">
        <h2>Create New Coffee</h2>
        <form
          onSubmit={handleFormSubmit}
          style={{
            marginBottom: "2rem",
            padding: "1rem",
            border: "1px solid #ccc",
          }}
        >
          <div>
            <label>Name: </label>
            <input
              type="text"
              name="name"
              value={formData.name}
              onChange={handleFormChange}
              required
            />
          </div>
          <div>
            <label>Roast: </label>
            <select
              name="roast"
              value={formData.roast}
              onChange={handleFormChange}
            >
              <option>Light</option>
              <option>Medium</option>
              <option>Dark</option>
            </select>
          </div>
          <div>
            <label>Image Path: </label>
            <input
              type="text"
              name="image"
              placeholder="/images/coffee.jpg"
              value={formData.image}
              onChange={handleFormChange}
            />
          </div>
          <div>
            <label>Region: </label>
            <input
              type="text"
              name="region"
              value={formData.region}
              onChange={handleFormChange}
              required
            />
          </div>
          <div>
            <label>Price: </label>
            <input
              type="number"
              step="0.01"
              name="price"
              value={formData.price}
              onChange={handleFormChange}
            />
          </div>
          <div>
            <label>Grind Unit (g): </label>
            <input
              type="number"
              name="grind_unit"
              value={formData.grind_unit}
              onChange={handleFormChange}
            />
          </div>
          <button type="submit">Create Coffee ☕</button>
        </form>

        <h2>Available Coffees</h2>
        {loading && <p>Loading coffees...</p>}
        {error && <p className="error">{error}</p>}

        {!loading && !error && coffees.length === 0 && (
          <p>No coffees yet. Create one from your backend API.</p>
        )}

        {!loading && !error && coffees.length > 0 && (
          <ul className="coffee-list">
            {coffees.map((coffee) => (
              <li key={coffee.id || coffee.name} className="coffee-card">
                <h3>{coffee.name}</h3>
                <img
                  style={{ width: "250px", height: "250px" }}
                  src={coffee.image || "/file.svg"}
                  alt={coffee.name || "Coffee image"}
                  onError={(e) => {
                    e.currentTarget.onerror = null;
                    e.currentTarget.src = "/file.svg";
                  }}
                />
                <p>
                  Origin: <strong>{coffee.region || "Unknown"}</strong>
                </p>
                <p>
                  Price: <strong>${coffee.price ?? "-"}</strong>
                </p>
                <button onClick={() => handleDelete(coffee)}>del🗑️</button>
              </li>
            ))}
          </ul>
        )}
      </main>
    </div>
  );
}
