import AddPrice from '@/components/AddPrice';
import Header from '@/components/Header'
import axios from "axios"
import {URL} from '@/Url'
import AddTable from '@/components/AddTable';
import UpdatePrice from '@/components/UpdatePrice';

const getData = async () =>{
  const response = await axios.get(`${URL}`)
  return response.data
}

export default async function HomePage() {
  const data = await getData()
  return (
    <main className="min-h-screen">
      <Header />
      <div className='buttons'>
        <div style={{gridColumn: '2/3'}}><AddTable /></div>
        <div style={{gridColumn: '3/4'}}><AddPrice data = {data} /></div>
        <div style={{gridColumn: '4/5'}}><UpdatePrice data = {data} /></div>
      </div>
      
      
    </main>
  );
}
