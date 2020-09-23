// Copyright © 2020 Weald Technology Trading.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package postgresql

import (
	"context"

	"github.com/wealdtech/chaind/services/chaindb"
)

// SetBlock sets a block.
func (s *Service) SetBlock(ctx context.Context, block *chaindb.Block) error {
	tx := s.tx(ctx)
	if tx == nil {
		return ErrNoTransaction
	}

	_, err := tx.Exec(ctx, `
      INSERT INTO t_blocks(f_slot
                          ,f_proposer_index
                          ,f_root
                          ,f_graffiti
                          ,f_randao_reveal
                          ,f_body_root
                          ,f_parent_root
                          ,f_state_root
                          ,f_eth1_block_hash
                          ,f_eth1_deposit_count
                          ,f_eth1_deposit_root
						  )
      VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
      ON CONFLICT (f_slot,f_root) DO
      UPDATE
      SET f_proposer_index = excluded.f_proposer_index
         ,f_root = excluded.f_root
         ,f_graffiti = excluded.f_graffiti
         ,f_randao_reveal = excluded.f_randao_reveal
         ,f_state_root = excluded.f_state_root
         ,f_eth1_block_hash = excluded.f_eth1_block_hash
         ,f_eth1_deposit_count = excluded.f_eth1_deposit_count
         ,f_eth1_deposit_root = excluded.f_eth1_deposit_root
	  `,
		block.Slot,
		block.ProposerIndex,
		block.Root,
		block.Graffiti,
		block.RANDAOReveal,
		block.BodyRoot,
		block.ParentRoot,
		block.StateRoot,
		block.ETH1BlockHash,
		block.ETH1DepositCount,
		block.ETH1DepositRoot,
	)

	return err
}