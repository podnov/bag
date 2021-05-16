import {
	GET_WALLET,
	GET_WALLET_ERROR
} from '../types'

import axios from 'axios'

const apiBaseUrl = '/bag/api/v1'

export const getWallet = (address) => async dispatch => {
	try{
		const res = await axios.get(`${apiBaseUrl}/accounts/${address}`)
		dispatch( {
			type: GET_WALLET,
			payload: res.data
		})
	}
	catch(e){
		dispatch( {
			type: GET_WALLET_ERROR,
			payload: console.log(e),
		})
	}
}
