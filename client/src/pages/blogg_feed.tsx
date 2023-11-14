import axios from "axios"

import styles from '../styles/blogg_feed.module.css'

import WebSocketComponent from '../components/websockets'
import { FormEvent } from "react"

export default function Stadschema() {
  const publishPost = (e: React.ChangeEvent<HTMLInputElement>)=>{
    e?.preventDefault()
    console.log(e.target["blogg_input"].value)
    axios.get('http://localhost:5001/api/send_websocket', {
    headers: {
        'Content-Type': 'application/json',
    },
    params: {
        blogg_input: e.target["blogg_input"].value,
    },
    })
    .then(function (response) {
        console.log(response);
    })
    .catch(function (error) {
        console.log(error);
    });
  }

  return (
    <div className={styles.main_container}>
      <div className={styles.feed_container}>
        <WebSocketComponent />
      </div>
      <form className={styles.editor_container} onSubmit={(data)=>{publishPost(data)}}>
        <h1>Blog Post Creator</h1>
        <input type='text' name="blogg_input" className={styles.editor} placeholder="Start writing your blog post..."></input>
        <button className={styles.btn_publish} type="submit">Publish</button>
      </form>
    </div>
  )
}
