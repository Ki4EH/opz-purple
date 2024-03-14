export interface Table{
    matrix_name: string
    location_id?: number
    microcategory_id?: number
    price?: string
    percent?: number
}

export interface Row{
    location_id: number
    microcategory_id: number
    price: string
}