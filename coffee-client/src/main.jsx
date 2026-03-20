import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import App from "./App.jsx";
import { Collections } from "./collection.jsx";
import "./styles/globals.css";
// router
import { BrowserRouter, Routes, Route } from "react-router-dom";
createRoot(document.getElementById("root")).render(
  <StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/home" element={<App />} />
        <Route path="/collections" element={<Collections />} />
      </Routes>
    </BrowserRouter>
  </StrictMode>,
);
