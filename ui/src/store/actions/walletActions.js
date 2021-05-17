import {
	GET_WALLET,
	GET_WALLET_ERROR,
	GET_WALLET_SUCCESS,
} from '../types'

import axios from 'axios'

const apiBaseUrl = 'bag/api/v1'
let apiHost = (process.env.NODE_ENV == 'production' ? 'http://api.cryptobag.podnov.com' : '')

export const getWallet = (address) => async dispatch => {
	dispatch({
		type: GET_WALLET
	})

	try{
		const res = await axios.get(`${apiHost}/${apiBaseUrl}/accounts/${address}`)
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
