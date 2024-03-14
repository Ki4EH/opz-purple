import React, {useState, useEffect} from 'react'
import { useForm } from 'react-hook-form'
import axios from 'axios'
import {URL} from "@/Url"
import styles from './AddPrice.module.css'

const UpdatePriceForm = ({props}: any)=> {
    const [matrix_names, setNames] = useState<any>([])
    const {register, handleSubmit} = useForm()

    const onSubmit = async (data: any) => {
        data ={
            "matrix_name": data.matrix_name,
            "microcategory_id": Number(data.microcategory_id),
            "location_id": Number(data.location_id),
            "price": Number(data.price),
            "percent": Number(data.percent)
        }
         axios.put(`${URL}/update`, data)
         .then(response => {
           alert('Строка успешно добавлена')
         });
         console.log(data);
         window.location.reload();
       };

    useEffect(() => {
        setNames(props)
      }, []);
  return (
    <div className={styles.container}>
        <form onSubmit={handleSubmit(onSubmit)}>
            <h2>Название таблицы:</h2>
            <select {...register("matrix_name")} required>
                {matrix_names.map((name:any)=><option key={name.id} value={name.id}>{name.id}</option>)}
            </select>
            <label>
                ID Микрокатегории:
                <input type='number' {...register("microcategory_id")}/>
            </label>
            <label>
                ID Локации:
                <input type='number' {...register("location_id")} />
            </label>
            <label>
                Цена:
                <input {...register("price")} placeholder='Для изменения цены'/>
            </label>
            <label>
                Процент:
                <input {...register("price")} placeholder='Для изменения цен на определенный процент'/>
            </label>
            <button type='submit'>Обновить</button>
        </form>
    </div>
  )
}

export default UpdatePriceForm