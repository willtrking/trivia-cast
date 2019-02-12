//
//  DoubleMediaBottomTitleViewController.swift
//  TriviaCast
//
//  Created by Brophy on 2/8/19.
//  Copyright © 2019 TriviaCast, Inc. All rights reserved.
//

import Foundation
import UIKit

class DoubleMediaBottomTitleViewController: UIViewController {
    
    private let leftViewContainer: UIView = UIView()
    private let rightViewContainer: UIView = UIView()
    private let titleLabel: UILabel = UILabel()
    
    private var rightMediaView: UIView?
    private var leftMediaView: UIView?
    
    override func viewDidLoad() {
        super.viewDidLoad()
        build()
        configure()
    }
    
    func setLeftMediaView(_ view: UIView?) {
        leftMediaView?.removeFromSuperview()
        
        guard let view = view else {
            return
        }
        leftViewContainer.addSubview(view)
        view.snp.makeConstraints { make in
            make.edges.equalTo(leftViewContainer)
        }
    }
    
    func setRightMediaView(_ view: UIView?) {
        rightMediaView?.removeFromSuperview()
        guard let view = view else {
            return
        }
        rightViewContainer.addSubview(view)
        view.snp.makeConstraints { make in
            make.edges.equalTo(rightViewContainer)
        }
    }
    
    func setTitle(_ title: String) {
        titleLabel.text = title
    }
    
    // MARK: - Private Interface
    
    private func build() {
        view.addSubview(leftViewContainer)
        view.addSubview(rightViewContainer)
        view.addSubview(titleLabel)
        
        titleLabel.font = TCFont.h2
        titleLabel.textAlignment = .center
        titleLabel.numberOfLines = 0
    }
    
    private func configure() {
        titleLabel.snp.makeConstraints { make in
            make.leading.trailing.bottom.equalTo(view)
        }
        leftViewContainer.snp.makeConstraints { make in
            make.leading.top.equalTo(view).offset(SpacingUnit.u4)
            make.trailing.equalTo(view.snp.centerX).multipliedBy(0.5).offset(-SpacingUnit.u2)
            make.bottom.equalTo(titleLabel.snp.top).offset(-SpacingUnit.u2)
        }
        rightViewContainer.snp.makeConstraints { make in
            make.top.equalTo(view).offset(SpacingUnit.u4)
            make.trailing.equalTo(view).offset(-SpacingUnit.u4)
            make.leading.equalTo(view.snp.centerX).multipliedBy(0.5).offset(SpacingUnit.u2)
            make.bottom.equalTo(titleLabel.snp.top).offset(-SpacingUnit.u2)
        }
    }
    
}
