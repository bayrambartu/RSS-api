import React from "react";
import UserForm from "./components/UserForm";
import FeedForm from "./components/FeedForm";
import SubscriptionForm from "./components/Subscription";
import PostList from "./components/PostList";

function App() {
    return (
        <div style={{ padding: "2rem" }}>
            <h1>RSS API Paneli</h1>
            <UserForm />
            <FeedForm />
            <SubscriptionForm/>
            <PostList/>
        </div>
    );
}

export default App;