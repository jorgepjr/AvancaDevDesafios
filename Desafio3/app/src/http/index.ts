import axios from 'axios'

const token = '4b72abc59236540db93d3ffac9e3efcaff59d80915303481c87b23a673346a77'

export const cryptoHttp = axios.create({
    baseURL: 'https://min-api.cryptocompare.com/data',
    headers: {
        authorization: `jjkey ${token}`
    }
})