import React, { Component } from 'react'
import {connect} from 'react-redux'
import './WalletInput.css';
import {getWallet} from '../../store/actions/walletActions'

class WalletInput extends Component {
	constructor(props) {
		super(props);

		this.state = {
			address: this.props.address
		}

		this.handleAddressChange = this.handleAddressChange.bind(this);
		this.handleFormSubmit = this.handleFormSubmit.bind(this);
	}

	handleAddressChange(event) {
		this.setState({address: event.target.value});
	}

	handleFormSubmit(event) {
		event.preventDefault();
		this.props.getWallet(this.state.address)
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

