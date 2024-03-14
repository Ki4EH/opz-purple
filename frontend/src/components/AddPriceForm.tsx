import React, {useState, useEffect} from 'react'
import { useForm } from 'react-hook-form'
import axios from 'axios'
import {URL} from "@/Url"
import styles from './AddPrice.module.css'


// const tables =[{"id": "baseline_matrix_1"},{"id":"baseline_matrix_2"},{"id":"baseline_matrix_3"},{"id":"discount_matrix_1"},{"id":"discount_matrix_2"},{"id":"discount_matrix_3"}]
const AddPriceForm = ({props}: any)=> {
    const {register, handleSubmit} = useForm()
    const [data, setData] = useState<any>([]);

    const onSubmit = async (data: any) => {
        data ={
            "matrix_name": data.matrix_name,
            "microcategory_id": Number(data.microcategory_id),
            "location_id": Number(data.location_id),
            "price": Number(data.price)
        }
         axios.post(`${URL}/add`, data)
         .then(response => {
           alert('Строка успешно добавлена')
         });
         console.log(data);
         window.location.reload();
       };

    useEffect(() => {
        console.log(props)
        setData(props)
      }, []);
  return (
    <div className={styles.container}>
        <form onSubmit={handleSubmit(onSubmit)}>
            <h2>Название таблицы:</h2>
            <select {...register("matrix_name")} required>
                {data.map((el:any)=><option key={el.id}>{el.id}</option>)}
            </select>
            <label>
                ID Микрокатегории:
                <input type='number' {...register("microcategory_id")} required/>
            </label>
            <label>
                ID Локации:
                <input type='number' {...register("location_id")} required/>
            </label>
            <label>
                Цена:
                <input {...register("price")} required/>
            </label>
            <button type='submit'>Добавить</button>
        </form>
    </div>
  )
}

export default AddPriceForm