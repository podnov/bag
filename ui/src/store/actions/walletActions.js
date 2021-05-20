import {
	GET_WALLET,
	GET_WALLET_ERROR,
	GET_WALLET_SUCCESS,
} from '../types'

import axios from 'axios'

const apiBaseUrl = 'bag/api/v1'
let productionApiHost = `${window.location.protocol}//api.cryptobag.podnov.com`;

let apiHost = (process.env.NODE_ENV === 'production' ? productionApiHost : window.location.href)

export const getWallet = (address) => async dispatch => {
	dispatch({
		type: GET_WALLET
	})

	let url = `${apiHost}/${apiBaseUrl}/accounts/${address}`

	try{
		const res = await axios.get(url)
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
