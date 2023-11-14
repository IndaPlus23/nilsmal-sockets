import React, { useState, useEffect } from 'react';
import styles from '../styles/blogg_feed.module.css'

const WebSocketComponent = () => {

    const [socket_state, set_socket_state] = useState<string[]>([]);

    useEffect(() => {
        const socket = new WebSocket('ws://localhost:5001/websocket'); // Replace with your WebSocket server URL

        socket.onopen = (event) => {
            console.log('WebSocket connection opened:', event);
        };

        socket.onmessage = (event) => {
            console.log('Received message from server:', event.data);
            set_socket_state((prevSocketState) => [event.data.toString(), ...prevSocketState]);

        };

        socket.onclose = (event) => {
            console.log('WebSocket connection closed:', event);
        };

        return () => {
            // Clean up the WebSocket connection when the component unmounts
            socket.close();
        };
    }, []);

    const display_sockets = () => {
        const socket_display_arr: JSX.Element[] = []
        socket_state.forEach(element => {
            socket_display_arr.push(
            <div className={styles.message_container}>
                <p>{element}</p>
            </div>
        )
        }); 
        return socket_display_arr.reverse()
    }

    return (
        display_sockets()
    );
};

export default WebSocketComponent;