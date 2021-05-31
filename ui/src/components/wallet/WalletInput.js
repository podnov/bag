import React, { Component } from 'react'
import {connect} from 'react-redux'
import './WalletInput.css';
import {getWallet} from '../../store/actions/walletActions'

const defaultAddress = "0xe453d625DfB485ab5f95Fa4CE689b2a06Ba8B2E7";
const localStorageWalletAddressKey = "walletAddress";

class WalletInput extends Component {

	constructor(props) {
		super(props);

		this.state = {
			address: this.props.address
		}

		this.handleAddressChange = this.handleAddressChange.bind(this);
		this.handleFormSubmit = this.handleFormSubmit.bind(this);
	}

	componentDidMount() {
		let initialAddress = localStorage.getItem(localStorageWalletAddressKey);

		if (!initialAddress) {
			initialAddress = defaultAddress;
		}

		this.setState({ address: initialAddress });
		this.props.getWallet(initialAddress);
	}

	handleAddressChange(event) {
		let address = event.target.value;

		localStorage.setItem(localStorageWalletAddressKey, address);
		this.setState({ address: address });
	}

	handleFormSubmit(event) {
		event.preventDefault();
		this.props.getWallet(this.state.address);
	}

	render() {
		return (
			<div className="walletInput">
				<form onSubmit={this.handleFormSubmit}>
					<label>
						Wallet Address:&nbsp;
						<input
							type="text"
							onChange={this.handleAddressChange}
							size="42"
							value={this.state.address}
							className="walletInputAddress"
							/>
					</label>
					&nbsp;
					<button type="submit">Go</button>
				</form>
			</div>
		)
	}
}

const mapStateToProps = (state) => ({address:state.address})

export default connect(mapStateToProps, {getWallet})(WalletInput)

