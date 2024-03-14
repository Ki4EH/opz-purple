"use client"
import React, { useState } from 'react';
import UpdatePriceForm from './UpdatePriceForm';
import styles from './AddPrice.module.css'

const UpdatePrice = ({data}: any) => {
    const [showOverlay, setShowOverlay] = useState(false);
    const handleClick = () => {
        setShowOverlay(!showOverlay);
    };
    return (
        
        <>
        <div className={styles.button} onClick={handleClick}>Изменить цену</div>
        {showOverlay && <UpdatePriceForm props={data} />}
        </>
    )
}

export default UpdatePrice