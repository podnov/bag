import {
	GET_WALLET,
	GET_WALLET_ERROR,
	GET_WALLET_SUCCESS
} from '../types'

const initialState = {
	address: null,
	hasError: false,
	isLoading: false,
	statistics: null
}

export default function reduce(state = initialState, action){
	switch(action.type){
		case GET_WALLET:
			return {
				...state,
				hasError: false,
				isLoading: true
			};
		case GET_WALLET_ERROR:
			return {
				...state,
				statistics: null,
				hasError: true,
				isLoading: false
			};
		case GET_WALLET_SUCCESS:
			return {
				...state,
				hasError: false,
				isLoading: false,
				statistics: action.payload
			};
		default:
			return state;
	}

}
