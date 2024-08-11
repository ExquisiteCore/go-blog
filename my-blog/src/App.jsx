import { useState } from 'react';

function App() {
  const [posts, setPosts] = useState([]);
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');

  const fetchPost = async (id) => {
    const response = await fetch(`http://localhost:8080/posts/${id}`);
    const data = await response.json();
    setPosts([data]);
  };

  const createPost = async () => {
    await fetch('http://localhost:8080/posts', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ title, content }),
    });
    setTitle('');
    setContent('');
  };

  return (
    <div>
      <h1>My Blog</h1>

      <input
        type="text"
        placeholder="Title"
        value={title}
        onChange={(e) => setTitle(e.target.value)}
      />
      <textarea
        placeholder="Content"
        value={content}
        onChange={(e) => setContent(e.target.value)}
      />
      <button onClick={createPost}>Create Post</button>

      <button onClick={() => fetchPost(1)}>Fetch Post 1</button>

      {posts.map((post, index) => (
        <div key={index}>
          <h2>{post.title}</h2>
          <div dangerouslySetInnerHTML={{ __html: post.content }} />
        </div>
      ))}
    </div>
  );
}

export default App;