import React, { useState } from "react";
import axios from "axios";

function SubscriptionForm() {
    const [userId, setUserId] = useState("");
    const [feedId, setFeedId] = useState("");
    const [apiKey, setApiKey] = useState("");
    const [status, setStatus] = useState("");

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            await axios.post(
                "http://localhost:3000/subscriptions",
                {
                    user_id: parseInt(userId),
                    feed_id: parseInt(feedId),
                },
                {
                    headers: {
                        Authorization: apiKey,
                    },
                }
            );
            setStatus("✅ Başarıyla abone olundu");
        } catch (err) {
            setStatus("❌ Abonelik başarısız");
            console.error(err);
        }
    };

    return (
        <div style={{ marginBottom: "2rem" }}>
            <h2>Feed Aboneliği</h2>
            <form onSubmit={handleSubmit}>
                <input
                    type="text"
                    placeholder="API Key"
                    value={apiKey}
                    onChange={(e) => setApiKey(e.target.value)}
                    required
                    style={{ marginRight: "1rem" }}
                />
                <input
                    type="number"
                    placeholder="Kullanıcı ID"
                    value={userId}
                    onChange={(e) => setUserId(e.target.value)}
                    required
                    style={{ marginRight: "1rem" }}
                />
                <input
                    type="number"
                    placeholder="Feed ID"
                    value={feedId}
                    onChange={(e) => setFeedId(e.target.value)}
                    required
                />
                <button type="submit" style={{ marginLeft: "1rem" }}>
                    Abone Ol
                </button>
            </form>
            {status && <p style={{ marginTop: "1rem" }}>{status}</p>}
        </div>
    );
}

export default SubscriptionForm;