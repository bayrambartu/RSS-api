import React, { useState } from "react";
import axios from "axios";

function FeedForm() {
    const [url, setUrl] = useState("");
    const [apiKey, setApiKey] = useState(""); // Kullanıcıdan alınacak
    const [status, setStatus] = useState("");

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            await axios.post(
                "http://localhost:3000/feeds",
                { url },
                {
                    headers: {
                        Authorization: apiKey,
                    },
                }
            );
            setStatus("✅ Feed başarıyla eklendi");
        } catch (err) {
            setStatus("❌ Feed eklenemedi");
            console.error(err);
        }
    };

    return (
        <div style={{ marginBottom: "2rem" }}>
            <h2>Feed Ekle</h2>
            <form onSubmit={handleSubmit}>
                <input
                    type="text"
                    placeholder="API Key"
                    value={apiKey}
                    onChange={(e) => setApiKey(e.target.value)}
                    required
                    style={{ marginRight: "1rem", width: "300px" }}
                />
                <input
                    type="text"
                    placeholder="RSS Feed URL"
                    value={url}
                    onChange={(e) => setUrl(e.target.value)}
                    required
                    style={{ marginRight: "1rem", width: "300px" }}
                />
                <button type="submit">Ekle</button>
            </form>
            {status && <p style={{ marginTop: "1rem" }}>{status}</p>}
        </div>
    );
}

export default FeedForm;