# new-cosmos-lms


./simd init --chain-id testnet myvalidator && ./simd add-genesis-account validator-key 1000000000stake && ./simd add-genesis-account $(./simd keys show mykey1 -a) 10000000000stake && ./simd gentx validator-key 1000000000stake --chain-id testnet && ./simd collect-gentxs && ./simd start