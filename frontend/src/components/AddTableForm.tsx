import React, {useState, useEffect} from 'react'
import { useForm } from 'react-hook-form'
import axios from 'axios'
import {URL} from "@/Url"
import styles from './AddPrice.module.css'
import { Row } from '@/types/Table.interface'

const AddTableForm = ()=> {
    const [matrix_names, setNames] = useState<any>([])
    const {register, handleSubmit} = useForm()

    const onSubmit = async (data: any) => {
         axios.post(`${URL}/add`, data)
         .then(response => {
           alert('Таблица успешно добавлена')
         });
         console.log(data);
         window.location.reload();
       };

       const [rows, setRows] = useState<Row[]>([]);
     
     const addRow = () => {
         setRows(prevRows => [...prevRows, {location_id: 0, microcategory_id: 0, price: ''  }]);
     };
  return (
    <div className={styles.container}>
        <form onSubmit={handleSubmit(onSubmit)}>
            <h2>Название таблицы:</h2>
            <input {...register("matrix_name")} required/>
            <h2>Строки:</h2>
            <div className={styles.rows}>
            <p>ID Микрокатегории</p>
            <p>ID Локации</p>
            <p>Цена</p>
            </div>
            
            {rows.map((row, index) => (
                <div className={styles.rows} key={index}>
                <input type='number'placeholder='ID Микрокатегории' {...register(`rows.${index}.microcategory_id`)} required/>
                <input type='number' placeholder='ID Локации' {...register(`rows.${index}.location_id`)} required/>
                <input placeholder='Цена' {...register(`rows.${index}.price`)} required/>
                </div>
            ))}
            <button type="button" onClick={addRow}>Добавить строку</button>
            <button type='submit'>Добавить</button>
        </form>
    </div>
  )
}

export default AddTableForm