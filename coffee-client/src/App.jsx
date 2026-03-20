import { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";

export default function App() {
  const [health, setHealth] = useState("Checking backend...");
  const [coffees, setCoffees] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

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
