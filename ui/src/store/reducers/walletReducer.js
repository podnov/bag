import {
	GET_WALLET,
	GET_WALLET_ERROR,
	GET_WALLET_SUCCESS
} from '../types'

const initialState = {
	address: null,
	isLoading: false,
	statistics: null
}

export default function reduce(state = initialState, action){
	switch(action.type){
		case GET_WALLET:
			console.log('get_wallet');
			return {
				...state,
				loading: true
			};
		case GET_WALLET_SUCCESS:
			console.log('get_wallet_success ' + action.payload);
			return {
				...state,
				statistics: action.payload,
				loading: false
			};
		default:
			return state;
	}

}
