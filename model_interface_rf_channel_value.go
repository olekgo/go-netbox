/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.7.0 (3.7)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// InterfaceRfChannelValue * `2.4g-1-2412-22` - 1 (2412 MHz) * `2.4g-2-2417-22` - 2 (2417 MHz) * `2.4g-3-2422-22` - 3 (2422 MHz) * `2.4g-4-2427-22` - 4 (2427 MHz) * `2.4g-5-2432-22` - 5 (2432 MHz) * `2.4g-6-2437-22` - 6 (2437 MHz) * `2.4g-7-2442-22` - 7 (2442 MHz) * `2.4g-8-2447-22` - 8 (2447 MHz) * `2.4g-9-2452-22` - 9 (2452 MHz) * `2.4g-10-2457-22` - 10 (2457 MHz) * `2.4g-11-2462-22` - 11 (2462 MHz) * `2.4g-12-2467-22` - 12 (2467 MHz) * `2.4g-13-2472-22` - 13 (2472 MHz) * `5g-32-5160-20` - 32 (5160/20 MHz) * `5g-34-5170-40` - 34 (5170/40 MHz) * `5g-36-5180-20` - 36 (5180/20 MHz) * `5g-38-5190-40` - 38 (5190/40 MHz) * `5g-40-5200-20` - 40 (5200/20 MHz) * `5g-42-5210-80` - 42 (5210/80 MHz) * `5g-44-5220-20` - 44 (5220/20 MHz) * `5g-46-5230-40` - 46 (5230/40 MHz) * `5g-48-5240-20` - 48 (5240/20 MHz) * `5g-50-5250-160` - 50 (5250/160 MHz) * `5g-52-5260-20` - 52 (5260/20 MHz) * `5g-54-5270-40` - 54 (5270/40 MHz) * `5g-56-5280-20` - 56 (5280/20 MHz) * `5g-58-5290-80` - 58 (5290/80 MHz) * `5g-60-5300-20` - 60 (5300/20 MHz) * `5g-62-5310-40` - 62 (5310/40 MHz) * `5g-64-5320-20` - 64 (5320/20 MHz) * `5g-100-5500-20` - 100 (5500/20 MHz) * `5g-102-5510-40` - 102 (5510/40 MHz) * `5g-104-5520-20` - 104 (5520/20 MHz) * `5g-106-5530-80` - 106 (5530/80 MHz) * `5g-108-5540-20` - 108 (5540/20 MHz) * `5g-110-5550-40` - 110 (5550/40 MHz) * `5g-112-5560-20` - 112 (5560/20 MHz) * `5g-114-5570-160` - 114 (5570/160 MHz) * `5g-116-5580-20` - 116 (5580/20 MHz) * `5g-118-5590-40` - 118 (5590/40 MHz) * `5g-120-5600-20` - 120 (5600/20 MHz) * `5g-122-5610-80` - 122 (5610/80 MHz) * `5g-124-5620-20` - 124 (5620/20 MHz) * `5g-126-5630-40` - 126 (5630/40 MHz) * `5g-128-5640-20` - 128 (5640/20 MHz) * `5g-132-5660-20` - 132 (5660/20 MHz) * `5g-134-5670-40` - 134 (5670/40 MHz) * `5g-136-5680-20` - 136 (5680/20 MHz) * `5g-138-5690-80` - 138 (5690/80 MHz) * `5g-140-5700-20` - 140 (5700/20 MHz) * `5g-142-5710-40` - 142 (5710/40 MHz) * `5g-144-5720-20` - 144 (5720/20 MHz) * `5g-149-5745-20` - 149 (5745/20 MHz) * `5g-151-5755-40` - 151 (5755/40 MHz) * `5g-153-5765-20` - 153 (5765/20 MHz) * `5g-155-5775-80` - 155 (5775/80 MHz) * `5g-157-5785-20` - 157 (5785/20 MHz) * `5g-159-5795-40` - 159 (5795/40 MHz) * `5g-161-5805-20` - 161 (5805/20 MHz) * `5g-163-5815-160` - 163 (5815/160 MHz) * `5g-165-5825-20` - 165 (5825/20 MHz) * `5g-167-5835-40` - 167 (5835/40 MHz) * `5g-169-5845-20` - 169 (5845/20 MHz) * `5g-171-5855-80` - 171 (5855/80 MHz) * `5g-173-5865-20` - 173 (5865/20 MHz) * `5g-175-5875-40` - 175 (5875/40 MHz) * `5g-177-5885-20` - 177 (5885/20 MHz) * `6g-1-5955-20` - 1 (5955/20 MHz) * `6g-3-5965-40` - 3 (5965/40 MHz) * `6g-5-5975-20` - 5 (5975/20 MHz) * `6g-7-5985-80` - 7 (5985/80 MHz) * `6g-9-5995-20` - 9 (5995/20 MHz) * `6g-11-6005-40` - 11 (6005/40 MHz) * `6g-13-6015-20` - 13 (6015/20 MHz) * `6g-15-6025-160` - 15 (6025/160 MHz) * `6g-17-6035-20` - 17 (6035/20 MHz) * `6g-19-6045-40` - 19 (6045/40 MHz) * `6g-21-6055-20` - 21 (6055/20 MHz) * `6g-23-6065-80` - 23 (6065/80 MHz) * `6g-25-6075-20` - 25 (6075/20 MHz) * `6g-27-6085-40` - 27 (6085/40 MHz) * `6g-29-6095-20` - 29 (6095/20 MHz) * `6g-31-6105-320` - 31 (6105/320 MHz) * `6g-33-6115-20` - 33 (6115/20 MHz) * `6g-35-6125-40` - 35 (6125/40 MHz) * `6g-37-6135-20` - 37 (6135/20 MHz) * `6g-39-6145-80` - 39 (6145/80 MHz) * `6g-41-6155-20` - 41 (6155/20 MHz) * `6g-43-6165-40` - 43 (6165/40 MHz) * `6g-45-6175-20` - 45 (6175/20 MHz) * `6g-47-6185-160` - 47 (6185/160 MHz) * `6g-49-6195-20` - 49 (6195/20 MHz) * `6g-51-6205-40` - 51 (6205/40 MHz) * `6g-53-6215-20` - 53 (6215/20 MHz) * `6g-55-6225-80` - 55 (6225/80 MHz) * `6g-57-6235-20` - 57 (6235/20 MHz) * `6g-59-6245-40` - 59 (6245/40 MHz) * `6g-61-6255-20` - 61 (6255/20 MHz) * `6g-65-6275-20` - 65 (6275/20 MHz) * `6g-67-6285-40` - 67 (6285/40 MHz) * `6g-69-6295-20` - 69 (6295/20 MHz) * `6g-71-6305-80` - 71 (6305/80 MHz) * `6g-73-6315-20` - 73 (6315/20 MHz) * `6g-75-6325-40` - 75 (6325/40 MHz) * `6g-77-6335-20` - 77 (6335/20 MHz) * `6g-79-6345-160` - 79 (6345/160 MHz) * `6g-81-6355-20` - 81 (6355/20 MHz) * `6g-83-6365-40` - 83 (6365/40 MHz) * `6g-85-6375-20` - 85 (6375/20 MHz) * `6g-87-6385-80` - 87 (6385/80 MHz) * `6g-89-6395-20` - 89 (6395/20 MHz) * `6g-91-6405-40` - 91 (6405/40 MHz) * `6g-93-6415-20` - 93 (6415/20 MHz) * `6g-95-6425-320` - 95 (6425/320 MHz) * `6g-97-6435-20` - 97 (6435/20 MHz) * `6g-99-6445-40` - 99 (6445/40 MHz) * `6g-101-6455-20` - 101 (6455/20 MHz) * `6g-103-6465-80` - 103 (6465/80 MHz) * `6g-105-6475-20` - 105 (6475/20 MHz) * `6g-107-6485-40` - 107 (6485/40 MHz) * `6g-109-6495-20` - 109 (6495/20 MHz) * `6g-111-6505-160` - 111 (6505/160 MHz) * `6g-113-6515-20` - 113 (6515/20 MHz) * `6g-115-6525-40` - 115 (6525/40 MHz) * `6g-117-6535-20` - 117 (6535/20 MHz) * `6g-119-6545-80` - 119 (6545/80 MHz) * `6g-121-6555-20` - 121 (6555/20 MHz) * `6g-123-6565-40` - 123 (6565/40 MHz) * `6g-125-6575-20` - 125 (6575/20 MHz) * `6g-129-6595-20` - 129 (6595/20 MHz) * `6g-131-6605-40` - 131 (6605/40 MHz) * `6g-133-6615-20` - 133 (6615/20 MHz) * `6g-135-6625-80` - 135 (6625/80 MHz) * `6g-137-6635-20` - 137 (6635/20 MHz) * `6g-139-6645-40` - 139 (6645/40 MHz) * `6g-141-6655-20` - 141 (6655/20 MHz) * `6g-143-6665-160` - 143 (6665/160 MHz) * `6g-145-6675-20` - 145 (6675/20 MHz) * `6g-147-6685-40` - 147 (6685/40 MHz) * `6g-149-6695-20` - 149 (6695/20 MHz) * `6g-151-6705-80` - 151 (6705/80 MHz) * `6g-153-6715-20` - 153 (6715/20 MHz) * `6g-155-6725-40` - 155 (6725/40 MHz) * `6g-157-6735-20` - 157 (6735/20 MHz) * `6g-159-6745-320` - 159 (6745/320 MHz) * `6g-161-6755-20` - 161 (6755/20 MHz) * `6g-163-6765-40` - 163 (6765/40 MHz) * `6g-165-6775-20` - 165 (6775/20 MHz) * `6g-167-6785-80` - 167 (6785/80 MHz) * `6g-169-6795-20` - 169 (6795/20 MHz) * `6g-171-6805-40` - 171 (6805/40 MHz) * `6g-173-6815-20` - 173 (6815/20 MHz) * `6g-175-6825-160` - 175 (6825/160 MHz) * `6g-177-6835-20` - 177 (6835/20 MHz) * `6g-179-6845-40` - 179 (6845/40 MHz) * `6g-181-6855-20` - 181 (6855/20 MHz) * `6g-183-6865-80` - 183 (6865/80 MHz) * `6g-185-6875-20` - 185 (6875/20 MHz) * `6g-187-6885-40` - 187 (6885/40 MHz) * `6g-189-6895-20` - 189 (6895/20 MHz) * `6g-193-6915-20` - 193 (6915/20 MHz) * `6g-195-6925-40` - 195 (6925/40 MHz) * `6g-197-6935-20` - 197 (6935/20 MHz) * `6g-199-6945-80` - 199 (6945/80 MHz) * `6g-201-6955-20` - 201 (6955/20 MHz) * `6g-203-6965-40` - 203 (6965/40 MHz) * `6g-205-6975-20` - 205 (6975/20 MHz) * `6g-207-6985-160` - 207 (6985/160 MHz) * `6g-209-6995-20` - 209 (6995/20 MHz) * `6g-211-7005-40` - 211 (7005/40 MHz) * `6g-213-7015-20` - 213 (7015/20 MHz) * `6g-215-7025-80` - 215 (7025/80 MHz) * `6g-217-7035-20` - 217 (7035/20 MHz) * `6g-219-7045-40` - 219 (7045/40 MHz) * `6g-221-7055-20` - 221 (7055/20 MHz) * `6g-225-7075-20` - 225 (7075/20 MHz) * `6g-227-7085-40` - 227 (7085/40 MHz) * `6g-229-7095-20` - 229 (7095/20 MHz) * `6g-233-7115-20` - 233 (7115/20 MHz) * `60g-1-58320-2160` - 1 (58.32/2.16 GHz) * `60g-2-60480-2160` - 2 (60.48/2.16 GHz) * `60g-3-62640-2160` - 3 (62.64/2.16 GHz) * `60g-4-64800-2160` - 4 (64.80/2.16 GHz) * `60g-5-66960-2160` - 5 (66.96/2.16 GHz) * `60g-6-69120-2160` - 6 (69.12/2.16 GHz) * `60g-9-59400-4320` - 9 (59.40/4.32 GHz) * `60g-10-61560-4320` - 10 (61.56/4.32 GHz) * `60g-11-63720-4320` - 11 (63.72/4.32 GHz) * `60g-12-65880-4320` - 12 (65.88/4.32 GHz) * `60g-13-68040-4320` - 13 (68.04/4.32 GHz) * `60g-17-60480-6480` - 17 (60.48/6.48 GHz) * `60g-18-62640-6480` - 18 (62.64/6.48 GHz) * `60g-19-64800-6480` - 19 (64.80/6.48 GHz) * `60g-20-66960-6480` - 20 (66.96/6.48 GHz) * `60g-25-61560-6480` - 25 (61.56/8.64 GHz) * `60g-26-63720-6480` - 26 (63.72/8.64 GHz) * `60g-27-65880-6480` - 27 (65.88/8.64 GHz)
type InterfaceRfChannelValue string

// List of Interface_rf_channel_value
const (
	INTERFACERFCHANNELVALUE__2_4G_1_2412_22    InterfaceRfChannelValue = "2.4g-1-2412-22"
	INTERFACERFCHANNELVALUE__2_4G_2_2417_22    InterfaceRfChannelValue = "2.4g-2-2417-22"
	INTERFACERFCHANNELVALUE__2_4G_3_2422_22    InterfaceRfChannelValue = "2.4g-3-2422-22"
	INTERFACERFCHANNELVALUE__2_4G_4_2427_22    InterfaceRfChannelValue = "2.4g-4-2427-22"
	INTERFACERFCHANNELVALUE__2_4G_5_2432_22    InterfaceRfChannelValue = "2.4g-5-2432-22"
	INTERFACERFCHANNELVALUE__2_4G_6_2437_22    InterfaceRfChannelValue = "2.4g-6-2437-22"
	INTERFACERFCHANNELVALUE__2_4G_7_2442_22    InterfaceRfChannelValue = "2.4g-7-2442-22"
	INTERFACERFCHANNELVALUE__2_4G_8_2447_22    InterfaceRfChannelValue = "2.4g-8-2447-22"
	INTERFACERFCHANNELVALUE__2_4G_9_2452_22    InterfaceRfChannelValue = "2.4g-9-2452-22"
	INTERFACERFCHANNELVALUE__2_4G_10_2457_22   InterfaceRfChannelValue = "2.4g-10-2457-22"
	INTERFACERFCHANNELVALUE__2_4G_11_2462_22   InterfaceRfChannelValue = "2.4g-11-2462-22"
	INTERFACERFCHANNELVALUE__2_4G_12_2467_22   InterfaceRfChannelValue = "2.4g-12-2467-22"
	INTERFACERFCHANNELVALUE__2_4G_13_2472_22   InterfaceRfChannelValue = "2.4g-13-2472-22"
	INTERFACERFCHANNELVALUE__5G_32_5160_20     InterfaceRfChannelValue = "5g-32-5160-20"
	INTERFACERFCHANNELVALUE__5G_34_5170_40     InterfaceRfChannelValue = "5g-34-5170-40"
	INTERFACERFCHANNELVALUE__5G_36_5180_20     InterfaceRfChannelValue = "5g-36-5180-20"
	INTERFACERFCHANNELVALUE__5G_38_5190_40     InterfaceRfChannelValue = "5g-38-5190-40"
	INTERFACERFCHANNELVALUE__5G_40_5200_20     InterfaceRfChannelValue = "5g-40-5200-20"
	INTERFACERFCHANNELVALUE__5G_42_5210_80     InterfaceRfChannelValue = "5g-42-5210-80"
	INTERFACERFCHANNELVALUE__5G_44_5220_20     InterfaceRfChannelValue = "5g-44-5220-20"
	INTERFACERFCHANNELVALUE__5G_46_5230_40     InterfaceRfChannelValue = "5g-46-5230-40"
	INTERFACERFCHANNELVALUE__5G_48_5240_20     InterfaceRfChannelValue = "5g-48-5240-20"
	INTERFACERFCHANNELVALUE__5G_50_5250_160    InterfaceRfChannelValue = "5g-50-5250-160"
	INTERFACERFCHANNELVALUE__5G_52_5260_20     InterfaceRfChannelValue = "5g-52-5260-20"
	INTERFACERFCHANNELVALUE__5G_54_5270_40     InterfaceRfChannelValue = "5g-54-5270-40"
	INTERFACERFCHANNELVALUE__5G_56_5280_20     InterfaceRfChannelValue = "5g-56-5280-20"
	INTERFACERFCHANNELVALUE__5G_58_5290_80     InterfaceRfChannelValue = "5g-58-5290-80"
	INTERFACERFCHANNELVALUE__5G_60_5300_20     InterfaceRfChannelValue = "5g-60-5300-20"
	INTERFACERFCHANNELVALUE__5G_62_5310_40     InterfaceRfChannelValue = "5g-62-5310-40"
	INTERFACERFCHANNELVALUE__5G_64_5320_20     InterfaceRfChannelValue = "5g-64-5320-20"
	INTERFACERFCHANNELVALUE__5G_100_5500_20    InterfaceRfChannelValue = "5g-100-5500-20"
	INTERFACERFCHANNELVALUE__5G_102_5510_40    InterfaceRfChannelValue = "5g-102-5510-40"
	INTERFACERFCHANNELVALUE__5G_104_5520_20    InterfaceRfChannelValue = "5g-104-5520-20"
	INTERFACERFCHANNELVALUE__5G_106_5530_80    InterfaceRfChannelValue = "5g-106-5530-80"
	INTERFACERFCHANNELVALUE__5G_108_5540_20    InterfaceRfChannelValue = "5g-108-5540-20"
	INTERFACERFCHANNELVALUE__5G_110_5550_40    InterfaceRfChannelValue = "5g-110-5550-40"
	INTERFACERFCHANNELVALUE__5G_112_5560_20    InterfaceRfChannelValue = "5g-112-5560-20"
	INTERFACERFCHANNELVALUE__5G_114_5570_160   InterfaceRfChannelValue = "5g-114-5570-160"
	INTERFACERFCHANNELVALUE__5G_116_5580_20    InterfaceRfChannelValue = "5g-116-5580-20"
	INTERFACERFCHANNELVALUE__5G_118_5590_40    InterfaceRfChannelValue = "5g-118-5590-40"
	INTERFACERFCHANNELVALUE__5G_120_5600_20    InterfaceRfChannelValue = "5g-120-5600-20"
	INTERFACERFCHANNELVALUE__5G_122_5610_80    InterfaceRfChannelValue = "5g-122-5610-80"
	INTERFACERFCHANNELVALUE__5G_124_5620_20    InterfaceRfChannelValue = "5g-124-5620-20"
	INTERFACERFCHANNELVALUE__5G_126_5630_40    InterfaceRfChannelValue = "5g-126-5630-40"
	INTERFACERFCHANNELVALUE__5G_128_5640_20    InterfaceRfChannelValue = "5g-128-5640-20"
	INTERFACERFCHANNELVALUE__5G_132_5660_20    InterfaceRfChannelValue = "5g-132-5660-20"
	INTERFACERFCHANNELVALUE__5G_134_5670_40    InterfaceRfChannelValue = "5g-134-5670-40"
	INTERFACERFCHANNELVALUE__5G_136_5680_20    InterfaceRfChannelValue = "5g-136-5680-20"
	INTERFACERFCHANNELVALUE__5G_138_5690_80    InterfaceRfChannelValue = "5g-138-5690-80"
	INTERFACERFCHANNELVALUE__5G_140_5700_20    InterfaceRfChannelValue = "5g-140-5700-20"
	INTERFACERFCHANNELVALUE__5G_142_5710_40    InterfaceRfChannelValue = "5g-142-5710-40"
	INTERFACERFCHANNELVALUE__5G_144_5720_20    InterfaceRfChannelValue = "5g-144-5720-20"
	INTERFACERFCHANNELVALUE__5G_149_5745_20    InterfaceRfChannelValue = "5g-149-5745-20"
	INTERFACERFCHANNELVALUE__5G_151_5755_40    InterfaceRfChannelValue = "5g-151-5755-40"
	INTERFACERFCHANNELVALUE__5G_153_5765_20    InterfaceRfChannelValue = "5g-153-5765-20"
	INTERFACERFCHANNELVALUE__5G_155_5775_80    InterfaceRfChannelValue = "5g-155-5775-80"
	INTERFACERFCHANNELVALUE__5G_157_5785_20    InterfaceRfChannelValue = "5g-157-5785-20"
	INTERFACERFCHANNELVALUE__5G_159_5795_40    InterfaceRfChannelValue = "5g-159-5795-40"
	INTERFACERFCHANNELVALUE__5G_161_5805_20    InterfaceRfChannelValue = "5g-161-5805-20"
	INTERFACERFCHANNELVALUE__5G_163_5815_160   InterfaceRfChannelValue = "5g-163-5815-160"
	INTERFACERFCHANNELVALUE__5G_165_5825_20    InterfaceRfChannelValue = "5g-165-5825-20"
	INTERFACERFCHANNELVALUE__5G_167_5835_40    InterfaceRfChannelValue = "5g-167-5835-40"
	INTERFACERFCHANNELVALUE__5G_169_5845_20    InterfaceRfChannelValue = "5g-169-5845-20"
	INTERFACERFCHANNELVALUE__5G_171_5855_80    InterfaceRfChannelValue = "5g-171-5855-80"
	INTERFACERFCHANNELVALUE__5G_173_5865_20    InterfaceRfChannelValue = "5g-173-5865-20"
	INTERFACERFCHANNELVALUE__5G_175_5875_40    InterfaceRfChannelValue = "5g-175-5875-40"
	INTERFACERFCHANNELVALUE__5G_177_5885_20    InterfaceRfChannelValue = "5g-177-5885-20"
	INTERFACERFCHANNELVALUE__6G_1_5955_20      InterfaceRfChannelValue = "6g-1-5955-20"
	INTERFACERFCHANNELVALUE__6G_3_5965_40      InterfaceRfChannelValue = "6g-3-5965-40"
	INTERFACERFCHANNELVALUE__6G_5_5975_20      InterfaceRfChannelValue = "6g-5-5975-20"
	INTERFACERFCHANNELVALUE__6G_7_5985_80      InterfaceRfChannelValue = "6g-7-5985-80"
	INTERFACERFCHANNELVALUE__6G_9_5995_20      InterfaceRfChannelValue = "6g-9-5995-20"
	INTERFACERFCHANNELVALUE__6G_11_6005_40     InterfaceRfChannelValue = "6g-11-6005-40"
	INTERFACERFCHANNELVALUE__6G_13_6015_20     InterfaceRfChannelValue = "6g-13-6015-20"
	INTERFACERFCHANNELVALUE__6G_15_6025_160    InterfaceRfChannelValue = "6g-15-6025-160"
	INTERFACERFCHANNELVALUE__6G_17_6035_20     InterfaceRfChannelValue = "6g-17-6035-20"
	INTERFACERFCHANNELVALUE__6G_19_6045_40     InterfaceRfChannelValue = "6g-19-6045-40"
	INTERFACERFCHANNELVALUE__6G_21_6055_20     InterfaceRfChannelValue = "6g-21-6055-20"
	INTERFACERFCHANNELVALUE__6G_23_6065_80     InterfaceRfChannelValue = "6g-23-6065-80"
	INTERFACERFCHANNELVALUE__6G_25_6075_20     InterfaceRfChannelValue = "6g-25-6075-20"
	INTERFACERFCHANNELVALUE__6G_27_6085_40     InterfaceRfChannelValue = "6g-27-6085-40"
	INTERFACERFCHANNELVALUE__6G_29_6095_20     InterfaceRfChannelValue = "6g-29-6095-20"
	INTERFACERFCHANNELVALUE__6G_31_6105_320    InterfaceRfChannelValue = "6g-31-6105-320"
	INTERFACERFCHANNELVALUE__6G_33_6115_20     InterfaceRfChannelValue = "6g-33-6115-20"
	INTERFACERFCHANNELVALUE__6G_35_6125_40     InterfaceRfChannelValue = "6g-35-6125-40"
	INTERFACERFCHANNELVALUE__6G_37_6135_20     InterfaceRfChannelValue = "6g-37-6135-20"
	INTERFACERFCHANNELVALUE__6G_39_6145_80     InterfaceRfChannelValue = "6g-39-6145-80"
	INTERFACERFCHANNELVALUE__6G_41_6155_20     InterfaceRfChannelValue = "6g-41-6155-20"
	INTERFACERFCHANNELVALUE__6G_43_6165_40     InterfaceRfChannelValue = "6g-43-6165-40"
	INTERFACERFCHANNELVALUE__6G_45_6175_20     InterfaceRfChannelValue = "6g-45-6175-20"
	INTERFACERFCHANNELVALUE__6G_47_6185_160    InterfaceRfChannelValue = "6g-47-6185-160"
	INTERFACERFCHANNELVALUE__6G_49_6195_20     InterfaceRfChannelValue = "6g-49-6195-20"
	INTERFACERFCHANNELVALUE__6G_51_6205_40     InterfaceRfChannelValue = "6g-51-6205-40"
	INTERFACERFCHANNELVALUE__6G_53_6215_20     InterfaceRfChannelValue = "6g-53-6215-20"
	INTERFACERFCHANNELVALUE__6G_55_6225_80     InterfaceRfChannelValue = "6g-55-6225-80"
	INTERFACERFCHANNELVALUE__6G_57_6235_20     InterfaceRfChannelValue = "6g-57-6235-20"
	INTERFACERFCHANNELVALUE__6G_59_6245_40     InterfaceRfChannelValue = "6g-59-6245-40"
	INTERFACERFCHANNELVALUE__6G_61_6255_20     InterfaceRfChannelValue = "6g-61-6255-20"
	INTERFACERFCHANNELVALUE__6G_65_6275_20     InterfaceRfChannelValue = "6g-65-6275-20"
	INTERFACERFCHANNELVALUE__6G_67_6285_40     InterfaceRfChannelValue = "6g-67-6285-40"
	INTERFACERFCHANNELVALUE__6G_69_6295_20     InterfaceRfChannelValue = "6g-69-6295-20"
	INTERFACERFCHANNELVALUE__6G_71_6305_80     InterfaceRfChannelValue = "6g-71-6305-80"
	INTERFACERFCHANNELVALUE__6G_73_6315_20     InterfaceRfChannelValue = "6g-73-6315-20"
	INTERFACERFCHANNELVALUE__6G_75_6325_40     InterfaceRfChannelValue = "6g-75-6325-40"
	INTERFACERFCHANNELVALUE__6G_77_6335_20     InterfaceRfChannelValue = "6g-77-6335-20"
	INTERFACERFCHANNELVALUE__6G_79_6345_160    InterfaceRfChannelValue = "6g-79-6345-160"
	INTERFACERFCHANNELVALUE__6G_81_6355_20     InterfaceRfChannelValue = "6g-81-6355-20"
	INTERFACERFCHANNELVALUE__6G_83_6365_40     InterfaceRfChannelValue = "6g-83-6365-40"
	INTERFACERFCHANNELVALUE__6G_85_6375_20     InterfaceRfChannelValue = "6g-85-6375-20"
	INTERFACERFCHANNELVALUE__6G_87_6385_80     InterfaceRfChannelValue = "6g-87-6385-80"
	INTERFACERFCHANNELVALUE__6G_89_6395_20     InterfaceRfChannelValue = "6g-89-6395-20"
	INTERFACERFCHANNELVALUE__6G_91_6405_40     InterfaceRfChannelValue = "6g-91-6405-40"
	INTERFACERFCHANNELVALUE__6G_93_6415_20     InterfaceRfChannelValue = "6g-93-6415-20"
	INTERFACERFCHANNELVALUE__6G_95_6425_320    InterfaceRfChannelValue = "6g-95-6425-320"
	INTERFACERFCHANNELVALUE__6G_97_6435_20     InterfaceRfChannelValue = "6g-97-6435-20"
	INTERFACERFCHANNELVALUE__6G_99_6445_40     InterfaceRfChannelValue = "6g-99-6445-40"
	INTERFACERFCHANNELVALUE__6G_101_6455_20    InterfaceRfChannelValue = "6g-101-6455-20"
	INTERFACERFCHANNELVALUE__6G_103_6465_80    InterfaceRfChannelValue = "6g-103-6465-80"
	INTERFACERFCHANNELVALUE__6G_105_6475_20    InterfaceRfChannelValue = "6g-105-6475-20"
	INTERFACERFCHANNELVALUE__6G_107_6485_40    InterfaceRfChannelValue = "6g-107-6485-40"
	INTERFACERFCHANNELVALUE__6G_109_6495_20    InterfaceRfChannelValue = "6g-109-6495-20"
	INTERFACERFCHANNELVALUE__6G_111_6505_160   InterfaceRfChannelValue = "6g-111-6505-160"
	INTERFACERFCHANNELVALUE__6G_113_6515_20    InterfaceRfChannelValue = "6g-113-6515-20"
	INTERFACERFCHANNELVALUE__6G_115_6525_40    InterfaceRfChannelValue = "6g-115-6525-40"
	INTERFACERFCHANNELVALUE__6G_117_6535_20    InterfaceRfChannelValue = "6g-117-6535-20"
	INTERFACERFCHANNELVALUE__6G_119_6545_80    InterfaceRfChannelValue = "6g-119-6545-80"
	INTERFACERFCHANNELVALUE__6G_121_6555_20    InterfaceRfChannelValue = "6g-121-6555-20"
	INTERFACERFCHANNELVALUE__6G_123_6565_40    InterfaceRfChannelValue = "6g-123-6565-40"
	INTERFACERFCHANNELVALUE__6G_125_6575_20    InterfaceRfChannelValue = "6g-125-6575-20"
	INTERFACERFCHANNELVALUE__6G_129_6595_20    InterfaceRfChannelValue = "6g-129-6595-20"
	INTERFACERFCHANNELVALUE__6G_131_6605_40    InterfaceRfChannelValue = "6g-131-6605-40"
	INTERFACERFCHANNELVALUE__6G_133_6615_20    InterfaceRfChannelValue = "6g-133-6615-20"
	INTERFACERFCHANNELVALUE__6G_135_6625_80    InterfaceRfChannelValue = "6g-135-6625-80"
	INTERFACERFCHANNELVALUE__6G_137_6635_20    InterfaceRfChannelValue = "6g-137-6635-20"
	INTERFACERFCHANNELVALUE__6G_139_6645_40    InterfaceRfChannelValue = "6g-139-6645-40"
	INTERFACERFCHANNELVALUE__6G_141_6655_20    InterfaceRfChannelValue = "6g-141-6655-20"
	INTERFACERFCHANNELVALUE__6G_143_6665_160   InterfaceRfChannelValue = "6g-143-6665-160"
	INTERFACERFCHANNELVALUE__6G_145_6675_20    InterfaceRfChannelValue = "6g-145-6675-20"
	INTERFACERFCHANNELVALUE__6G_147_6685_40    InterfaceRfChannelValue = "6g-147-6685-40"
	INTERFACERFCHANNELVALUE__6G_149_6695_20    InterfaceRfChannelValue = "6g-149-6695-20"
	INTERFACERFCHANNELVALUE__6G_151_6705_80    InterfaceRfChannelValue = "6g-151-6705-80"
	INTERFACERFCHANNELVALUE__6G_153_6715_20    InterfaceRfChannelValue = "6g-153-6715-20"
	INTERFACERFCHANNELVALUE__6G_155_6725_40    InterfaceRfChannelValue = "6g-155-6725-40"
	INTERFACERFCHANNELVALUE__6G_157_6735_20    InterfaceRfChannelValue = "6g-157-6735-20"
	INTERFACERFCHANNELVALUE__6G_159_6745_320   InterfaceRfChannelValue = "6g-159-6745-320"
	INTERFACERFCHANNELVALUE__6G_161_6755_20    InterfaceRfChannelValue = "6g-161-6755-20"
	INTERFACERFCHANNELVALUE__6G_163_6765_40    InterfaceRfChannelValue = "6g-163-6765-40"
	INTERFACERFCHANNELVALUE__6G_165_6775_20    InterfaceRfChannelValue = "6g-165-6775-20"
	INTERFACERFCHANNELVALUE__6G_167_6785_80    InterfaceRfChannelValue = "6g-167-6785-80"
	INTERFACERFCHANNELVALUE__6G_169_6795_20    InterfaceRfChannelValue = "6g-169-6795-20"
	INTERFACERFCHANNELVALUE__6G_171_6805_40    InterfaceRfChannelValue = "6g-171-6805-40"
	INTERFACERFCHANNELVALUE__6G_173_6815_20    InterfaceRfChannelValue = "6g-173-6815-20"
	INTERFACERFCHANNELVALUE__6G_175_6825_160   InterfaceRfChannelValue = "6g-175-6825-160"
	INTERFACERFCHANNELVALUE__6G_177_6835_20    InterfaceRfChannelValue = "6g-177-6835-20"
	INTERFACERFCHANNELVALUE__6G_179_6845_40    InterfaceRfChannelValue = "6g-179-6845-40"
	INTERFACERFCHANNELVALUE__6G_181_6855_20    InterfaceRfChannelValue = "6g-181-6855-20"
	INTERFACERFCHANNELVALUE__6G_183_6865_80    InterfaceRfChannelValue = "6g-183-6865-80"
	INTERFACERFCHANNELVALUE__6G_185_6875_20    InterfaceRfChannelValue = "6g-185-6875-20"
	INTERFACERFCHANNELVALUE__6G_187_6885_40    InterfaceRfChannelValue = "6g-187-6885-40"
	INTERFACERFCHANNELVALUE__6G_189_6895_20    InterfaceRfChannelValue = "6g-189-6895-20"
	INTERFACERFCHANNELVALUE__6G_193_6915_20    InterfaceRfChannelValue = "6g-193-6915-20"
	INTERFACERFCHANNELVALUE__6G_195_6925_40    InterfaceRfChannelValue = "6g-195-6925-40"
	INTERFACERFCHANNELVALUE__6G_197_6935_20    InterfaceRfChannelValue = "6g-197-6935-20"
	INTERFACERFCHANNELVALUE__6G_199_6945_80    InterfaceRfChannelValue = "6g-199-6945-80"
	INTERFACERFCHANNELVALUE__6G_201_6955_20    InterfaceRfChannelValue = "6g-201-6955-20"
	INTERFACERFCHANNELVALUE__6G_203_6965_40    InterfaceRfChannelValue = "6g-203-6965-40"
	INTERFACERFCHANNELVALUE__6G_205_6975_20    InterfaceRfChannelValue = "6g-205-6975-20"
	INTERFACERFCHANNELVALUE__6G_207_6985_160   InterfaceRfChannelValue = "6g-207-6985-160"
	INTERFACERFCHANNELVALUE__6G_209_6995_20    InterfaceRfChannelValue = "6g-209-6995-20"
	INTERFACERFCHANNELVALUE__6G_211_7005_40    InterfaceRfChannelValue = "6g-211-7005-40"
	INTERFACERFCHANNELVALUE__6G_213_7015_20    InterfaceRfChannelValue = "6g-213-7015-20"
	INTERFACERFCHANNELVALUE__6G_215_7025_80    InterfaceRfChannelValue = "6g-215-7025-80"
	INTERFACERFCHANNELVALUE__6G_217_7035_20    InterfaceRfChannelValue = "6g-217-7035-20"
	INTERFACERFCHANNELVALUE__6G_219_7045_40    InterfaceRfChannelValue = "6g-219-7045-40"
	INTERFACERFCHANNELVALUE__6G_221_7055_20    InterfaceRfChannelValue = "6g-221-7055-20"
	INTERFACERFCHANNELVALUE__6G_225_7075_20    InterfaceRfChannelValue = "6g-225-7075-20"
	INTERFACERFCHANNELVALUE__6G_227_7085_40    InterfaceRfChannelValue = "6g-227-7085-40"
	INTERFACERFCHANNELVALUE__6G_229_7095_20    InterfaceRfChannelValue = "6g-229-7095-20"
	INTERFACERFCHANNELVALUE__6G_233_7115_20    InterfaceRfChannelValue = "6g-233-7115-20"
	INTERFACERFCHANNELVALUE__60G_1_58320_2160  InterfaceRfChannelValue = "60g-1-58320-2160"
	INTERFACERFCHANNELVALUE__60G_2_60480_2160  InterfaceRfChannelValue = "60g-2-60480-2160"
	INTERFACERFCHANNELVALUE__60G_3_62640_2160  InterfaceRfChannelValue = "60g-3-62640-2160"
	INTERFACERFCHANNELVALUE__60G_4_64800_2160  InterfaceRfChannelValue = "60g-4-64800-2160"
	INTERFACERFCHANNELVALUE__60G_5_66960_2160  InterfaceRfChannelValue = "60g-5-66960-2160"
	INTERFACERFCHANNELVALUE__60G_6_69120_2160  InterfaceRfChannelValue = "60g-6-69120-2160"
	INTERFACERFCHANNELVALUE__60G_9_59400_4320  InterfaceRfChannelValue = "60g-9-59400-4320"
	INTERFACERFCHANNELVALUE__60G_10_61560_4320 InterfaceRfChannelValue = "60g-10-61560-4320"
	INTERFACERFCHANNELVALUE__60G_11_63720_4320 InterfaceRfChannelValue = "60g-11-63720-4320"
	INTERFACERFCHANNELVALUE__60G_12_65880_4320 InterfaceRfChannelValue = "60g-12-65880-4320"
	INTERFACERFCHANNELVALUE__60G_13_68040_4320 InterfaceRfChannelValue = "60g-13-68040-4320"
	INTERFACERFCHANNELVALUE__60G_17_60480_6480 InterfaceRfChannelValue = "60g-17-60480-6480"
	INTERFACERFCHANNELVALUE__60G_18_62640_6480 InterfaceRfChannelValue = "60g-18-62640-6480"
	INTERFACERFCHANNELVALUE__60G_19_64800_6480 InterfaceRfChannelValue = "60g-19-64800-6480"
	INTERFACERFCHANNELVALUE__60G_20_66960_6480 InterfaceRfChannelValue = "60g-20-66960-6480"
	INTERFACERFCHANNELVALUE__60G_25_61560_6480 InterfaceRfChannelValue = "60g-25-61560-6480"
	INTERFACERFCHANNELVALUE__60G_26_63720_6480 InterfaceRfChannelValue = "60g-26-63720-6480"
	INTERFACERFCHANNELVALUE__60G_27_65880_6480 InterfaceRfChannelValue = "60g-27-65880-6480"
	INTERFACERFCHANNELVALUE_EMPTY              InterfaceRfChannelValue = ""
)

// All allowed values of InterfaceRfChannelValue enum
var AllowedInterfaceRfChannelValueEnumValues = []InterfaceRfChannelValue{
	"2.4g-1-2412-22",
	"2.4g-2-2417-22",
	"2.4g-3-2422-22",
	"2.4g-4-2427-22",
	"2.4g-5-2432-22",
	"2.4g-6-2437-22",
	"2.4g-7-2442-22",
	"2.4g-8-2447-22",
	"2.4g-9-2452-22",
	"2.4g-10-2457-22",
	"2.4g-11-2462-22",
	"2.4g-12-2467-22",
	"2.4g-13-2472-22",
	"5g-32-5160-20",
	"5g-34-5170-40",
	"5g-36-5180-20",
	"5g-38-5190-40",
	"5g-40-5200-20",
	"5g-42-5210-80",
	"5g-44-5220-20",
	"5g-46-5230-40",
	"5g-48-5240-20",
	"5g-50-5250-160",
	"5g-52-5260-20",
	"5g-54-5270-40",
	"5g-56-5280-20",
	"5g-58-5290-80",
	"5g-60-5300-20",
	"5g-62-5310-40",
	"5g-64-5320-20",
	"5g-100-5500-20",
	"5g-102-5510-40",
	"5g-104-5520-20",
	"5g-106-5530-80",
	"5g-108-5540-20",
	"5g-110-5550-40",
	"5g-112-5560-20",
	"5g-114-5570-160",
	"5g-116-5580-20",
	"5g-118-5590-40",
	"5g-120-5600-20",
	"5g-122-5610-80",
	"5g-124-5620-20",
	"5g-126-5630-40",
	"5g-128-5640-20",
	"5g-132-5660-20",
	"5g-134-5670-40",
	"5g-136-5680-20",
	"5g-138-5690-80",
	"5g-140-5700-20",
	"5g-142-5710-40",
	"5g-144-5720-20",
	"5g-149-5745-20",
	"5g-151-5755-40",
	"5g-153-5765-20",
	"5g-155-5775-80",
	"5g-157-5785-20",
	"5g-159-5795-40",
	"5g-161-5805-20",
	"5g-163-5815-160",
	"5g-165-5825-20",
	"5g-167-5835-40",
	"5g-169-5845-20",
	"5g-171-5855-80",
	"5g-173-5865-20",
	"5g-175-5875-40",
	"5g-177-5885-20",
	"6g-1-5955-20",
	"6g-3-5965-40",
	"6g-5-5975-20",
	"6g-7-5985-80",
	"6g-9-5995-20",
	"6g-11-6005-40",
	"6g-13-6015-20",
	"6g-15-6025-160",
	"6g-17-6035-20",
	"6g-19-6045-40",
	"6g-21-6055-20",
	"6g-23-6065-80",
	"6g-25-6075-20",
	"6g-27-6085-40",
	"6g-29-6095-20",
	"6g-31-6105-320",
	"6g-33-6115-20",
	"6g-35-6125-40",
	"6g-37-6135-20",
	"6g-39-6145-80",
	"6g-41-6155-20",
	"6g-43-6165-40",
	"6g-45-6175-20",
	"6g-47-6185-160",
	"6g-49-6195-20",
	"6g-51-6205-40",
	"6g-53-6215-20",
	"6g-55-6225-80",
	"6g-57-6235-20",
	"6g-59-6245-40",
	"6g-61-6255-20",
	"6g-65-6275-20",
	"6g-67-6285-40",
	"6g-69-6295-20",
	"6g-71-6305-80",
	"6g-73-6315-20",
	"6g-75-6325-40",
	"6g-77-6335-20",
	"6g-79-6345-160",
	"6g-81-6355-20",
	"6g-83-6365-40",
	"6g-85-6375-20",
	"6g-87-6385-80",
	"6g-89-6395-20",
	"6g-91-6405-40",
	"6g-93-6415-20",
	"6g-95-6425-320",
	"6g-97-6435-20",
	"6g-99-6445-40",
	"6g-101-6455-20",
	"6g-103-6465-80",
	"6g-105-6475-20",
	"6g-107-6485-40",
	"6g-109-6495-20",
	"6g-111-6505-160",
	"6g-113-6515-20",
	"6g-115-6525-40",
	"6g-117-6535-20",
	"6g-119-6545-80",
	"6g-121-6555-20",
	"6g-123-6565-40",
	"6g-125-6575-20",
	"6g-129-6595-20",
	"6g-131-6605-40",
	"6g-133-6615-20",
	"6g-135-6625-80",
	"6g-137-6635-20",
	"6g-139-6645-40",
	"6g-141-6655-20",
	"6g-143-6665-160",
	"6g-145-6675-20",
	"6g-147-6685-40",
	"6g-149-6695-20",
	"6g-151-6705-80",
	"6g-153-6715-20",
	"6g-155-6725-40",
	"6g-157-6735-20",
	"6g-159-6745-320",
	"6g-161-6755-20",
	"6g-163-6765-40",
	"6g-165-6775-20",
	"6g-167-6785-80",
	"6g-169-6795-20",
	"6g-171-6805-40",
	"6g-173-6815-20",
	"6g-175-6825-160",
	"6g-177-6835-20",
	"6g-179-6845-40",
	"6g-181-6855-20",
	"6g-183-6865-80",
	"6g-185-6875-20",
	"6g-187-6885-40",
	"6g-189-6895-20",
	"6g-193-6915-20",
	"6g-195-6925-40",
	"6g-197-6935-20",
	"6g-199-6945-80",
	"6g-201-6955-20",
	"6g-203-6965-40",
	"6g-205-6975-20",
	"6g-207-6985-160",
	"6g-209-6995-20",
	"6g-211-7005-40",
	"6g-213-7015-20",
	"6g-215-7025-80",
	"6g-217-7035-20",
	"6g-219-7045-40",
	"6g-221-7055-20",
	"6g-225-7075-20",
	"6g-227-7085-40",
	"6g-229-7095-20",
	"6g-233-7115-20",
	"60g-1-58320-2160",
	"60g-2-60480-2160",
	"60g-3-62640-2160",
	"60g-4-64800-2160",
	"60g-5-66960-2160",
	"60g-6-69120-2160",
	"60g-9-59400-4320",
	"60g-10-61560-4320",
	"60g-11-63720-4320",
	"60g-12-65880-4320",
	"60g-13-68040-4320",
	"60g-17-60480-6480",
	"60g-18-62640-6480",
	"60g-19-64800-6480",
	"60g-20-66960-6480",
	"60g-25-61560-6480",
	"60g-26-63720-6480",
	"60g-27-65880-6480",
	"",
}

func (v *InterfaceRfChannelValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InterfaceRfChannelValue(value)
	for _, existing := range AllowedInterfaceRfChannelValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InterfaceRfChannelValue", value)
}

// NewInterfaceRfChannelValueFromValue returns a pointer to a valid InterfaceRfChannelValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInterfaceRfChannelValueFromValue(v string) (*InterfaceRfChannelValue, error) {
	ev := InterfaceRfChannelValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InterfaceRfChannelValue: valid values are %v", v, AllowedInterfaceRfChannelValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InterfaceRfChannelValue) IsValid() bool {
	for _, existing := range AllowedInterfaceRfChannelValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Interface_rf_channel_value value
func (v InterfaceRfChannelValue) Ptr() *InterfaceRfChannelValue {
	return &v
}

type NullableInterfaceRfChannelValue struct {
	value *InterfaceRfChannelValue
	isSet bool
}

func (v NullableInterfaceRfChannelValue) Get() *InterfaceRfChannelValue {
	return v.value
}

func (v *NullableInterfaceRfChannelValue) Set(val *InterfaceRfChannelValue) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceRfChannelValue) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceRfChannelValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceRfChannelValue(val *InterfaceRfChannelValue) *NullableInterfaceRfChannelValue {
	return &NullableInterfaceRfChannelValue{value: val, isSet: true}
}

func (v NullableInterfaceRfChannelValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceRfChannelValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
