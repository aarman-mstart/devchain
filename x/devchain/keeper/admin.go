package keeper
import "context"
const DefaultAdminAddress = "cosmos1l9zg8wns3nxltycfyljp328c2wulwqct8mq62k"

func(k Keeper) GetAdminAddress(ctx context.Context) (string, error) {
	params, err := k.Params.Get(ctx)
	
	if err != nil {
		return "", err
	}
	if params.Admin != "" {
		return params.Admin, nil
	}
	return DefaultAdminAddress, nil
}