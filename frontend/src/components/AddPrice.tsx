"use client"
import React, { useState } from 'react';
import AddPriceForm from './AddPriceForm';
import styles from './AddPrice.module.css'

const AddPrice = ({data}: any) => {
    const [showOverlay, setShowOverlay] = useState(false);
    const handleClick = () => {
        setShowOverlay(!showOverlay);
    };
    return (
        
        <>
        <div className={styles.button} onClick={handleClick}>Добавить цену</div>
        {showOverlay && <AddPriceForm props={data} />}
        </>
    )
}

export default AddPrice