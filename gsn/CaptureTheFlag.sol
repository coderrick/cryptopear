/**
 * SPDX-License-Identifier:MIT
 */
pragma solidity >=0.4.22 <0.7.0;



import "./BaseRelayRecipient.sol";
import "./IKnowForwarderAddress.sol";

contract CaptureTheFlag is BaseRelayRecipient, IKnowForwarderAddress {
	address flagHolder = address(0);
	event FlagCaptured(address _from, address _to);

	constructor(address forwarder) public {
		setTrustedForwarder(forwarder);
	}

	function getTrustedForwarder() public override view returns(address) {
		return trustedForwarder;
	}

	function setTrustedForwarder(address forwarder) internal {
		trustedForwarder = forwarder;
	}


	function captureFlag() external {
		address previous = flagHolder;

                // The real sender. If you are using GSNv2, this
                // is not the same as msg.sender.
		flagHolder = _msgSender();  

		emit FlagCaptured(previous, flagHolder); 
	}

	function versionRecipient() external virtual view 
	override returns (string memory) {
		return "1.0";
	}

}