import React from 'react'
import styles from './Header.module.css'
import Image from 'next/image'
const Header = () => {
  return (
    <div className={styles.container}>
        <div className={styles.logo}>
        <Image 
        src="/Avito_logo.svg.png"
        width={40}
        height={40}
        alt='' />
        <div className={styles.logo_title}>AvitoPrice</div>
        </div>
        
    </div>
  )
}

export default Header