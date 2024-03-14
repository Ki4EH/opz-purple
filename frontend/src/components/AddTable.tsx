"use client"
import React, { useState } from 'react';
import AddTableForm from './AddTableForm';
import styles from './AddPrice.module.css'

const AddTable = ({data}: any) => {
    const [showOverlay, setShowOverlay] = useState(false);
    const handleClick = () => {
        setShowOverlay(!showOverlay);
    };
    return (
        
        <>
        <div className={styles.button} onClick={handleClick}>Создать таблицу</div>
        {showOverlay && <AddTableForm />}
        </>
    )
}

export default AddTable