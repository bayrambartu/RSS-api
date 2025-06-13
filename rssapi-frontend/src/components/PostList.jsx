import React, { useState } from "react";
import axios from "axios";

function PostList() {
    const [apiKey, setApiKey] = useState("");
    const [posts, setPosts] = useState([]);
    const [status, setStatus] = useState("");

    const fetchPosts = async () => {
        try {
            const res = await axios.get("http://localhost:3000/posts", {
                headers: {
                    Authorization: apiKey,
                },
            });
            setPosts(res.data);
            setStatus("✅ Gönderiler başarıyla yüklendi");
        } catch (err) {
            setStatus("❌ Gönderiler alınamadı");
            console.error(err);
        }
    };

    return (
        <div style={{ marginBottom: "2rem" }}>
            <h2>Gönderileri Görüntüle</h2>
            <input
                type="text"
                placeholder="API Key"
                value={apiKey}
                onChange={(e) => setApiKey(e.target.value)}
                required
                style={{ marginRight: "1rem", width: "300px" }}
            />
            <button onClick={fetchPosts}>Getir</button>
            {status && <p style={{ marginTop: "1rem" }}>{status}</p>}

            <ul>
                {posts.map((post) => (
                    <li key={post.id} style={{ marginBottom: "1rem" }}>
                        <a href={post.url} target="_blank" rel="noopener noreferrer">
                            {post.title}
                        </a>
                        <br />
                        <small>{new Date(post.published_at).toLocaleString()}</small>
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default PostList;