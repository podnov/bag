import React, { Component } from 'react'
import {connect} from 'react-redux'
import {getWallet} from '../../store/actions/walletActions'

class WalletInput extends Component {
	constructor(props) {
		super(props);

		this.state = {
			address: this.props.address
		}

		this.handleAddressChange = this.handleAddressChange.bind(this);
		this.handleGoClick = this.handleGoClick.bind(this);
	}

	handleAddressChange(event) {
		this.setState({address: event.target.value});
	}

	handleGoClick(event) {
		this.props.getWallet(this.state.address)
	}

	render() {
		return (
			<div className="walletInput">
				<label>
					Wallet Address:
					<input
						type="text"
						onChange={this.handleAddressChange}
						value={this.state.address}
						/>
					<button
						onClick={this.handleGoClick}
						>
						Go
					</button>
				</label>
			</div>
		)
	}
}

const mapStateToProps = (state) => ({address:state.address})

export default connect(mapStateToProps, {getWallet})(WalletInput)

