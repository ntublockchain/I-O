package fabric_asset

// func VerifyKnownSignatures(hash []byte, sigs [][]byte, addrs []common.Address) bool {
// 	set := make(map[string]bool, len(addrs))
// 	for _, addr := range addrs {
// 		set[addr.Hex()] = true
// 	}

// 	for _, sig := range sigs {
// 		addr, err := VerifySignature(hash, sig)
// 		if err != nil {
// 			return false
// 		}
// 		if !set[addr.Hex()] {
// 			return false
// 		}
// 	}
// 	return true
// }
