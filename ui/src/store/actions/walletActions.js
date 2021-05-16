import {
	GET_WALLET,
	GET_WALLET_ERROR,
	GET_WALLET_SUCCESS,
} from '../types'

import axios from 'axios'

const apiBaseUrl = '/bag/api/v1'

export const getWallet = (address) => async dispatch => {
	dispatch({
		type: GET_WALLET
	})

	try{
		const res = await axios.get(`${apiBaseUrl}/accounts/${address}`)
		dispatch({
			type: GET_WALLET_SUCCESS,
			payload: res.data
		})
	} catch(e) {
		dispatch({
			type: GET_WALLET_ERROR,
			payload: console.log(e),
		})
	}
}
