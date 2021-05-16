import {
	GET_WALLET,
	GET_WALLET_ERROR
} from '../types'

const initialState = {
	walletAddress:'',
	wallet:[],
	loading:false
}

export default function(state = initialState, action){
	switch(action.type){
		case GET_WALLET:
		return {
			...state,
			wallet:action.payload,
			loading:false

		}
		default: return state
	}

}
