import React, { useState } from "react";
import axios from "axios";

function UserForm() {
    const [name, setName] = useState("");
    const [email, setEmail] = useState("");
    const [apiKey, setApiKey] = useState("");

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const res = await axios.post("http://localhost:3000/users", {
                name,
                email,
            });
            setApiKey(res.data.api_key);
        } catch (err) {
            alert("Kullanıcı oluşturulamadı ❌");
            console.error(err);
        }
    };

    return (
        <div style={{ marginBottom: "2rem" }}>
            <h2>Kullanıcı Oluştur</h2>
            <form onSubmit={handleSubmit}>
                <input
                    type="text"
                    placeholder="Ad"
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                    required
                    style={{ marginRight: "1rem" }}
                />
                <input
                    type="email"
                    placeholder="E-posta"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    required
                />
                <button type="submit" style={{ marginLeft: "1rem" }}>
                    Oluştur
                </button>
            </form>
            {apiKey && (
                <div style={{ marginTop: "1rem", color: "green" }}>
                    ✅ API Key: <strong>{apiKey}</strong>
                </div>
            )}
        </div>
    );
}

export default UserForm;