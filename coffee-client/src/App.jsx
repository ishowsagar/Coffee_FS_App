import { useEffect, useState } from "react";

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

  return (
    <div className="page">
      <header className="hero">
        <h1>Coffee Hub</h1>
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
              </li>
            ))}
          </ul>
        )}
      </main>
    </div>
  );
}
